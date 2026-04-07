import { compile } from 'svelte/compiler';
import { render } from 'svelte/server';
import chokidar from 'chokidar';
import esbuild from 'esbuild';
import { existsSync, readdirSync, readFileSync, rmSync, statSync, writeFileSync } from 'fs';
import { basename, dirname, join, relative, resolve } from 'path';
import { fileURLToPath } from 'url';
import sveltePlugin from 'esbuild-svelte';
import { sum } from 'lodash-es';
import { parse, serialize } from 'parse5';
import notifier from 'node-notifier';
import svelteConfig from './svelte.config.js';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

process.on('uncaughtException', error => {
  console.error(error);
  notifier.notify({
    title: 'Error occurs',
    message: `${error}`,
  });
});

process.on('unhandledRejection', error => {
  console.error(error);
  notifier.notify({
    title: 'Unhandled rejection',
    message: `${error}`,
  });
});

const [watch, serve, minify, debug, logVars] = ['--watch', '--serve', '--minify', '--debug', '--log-vars'].map(s =>
  process.argv.includes(s)
);

const ignorePath = new Set([
  'node_modules',
  '.vscode',
  '.idea',
  '.git',
  '.gitignore',
  '.ssr_temp',
  'build.js',
  'package-lock.json',
  'package.json',
  'README.md',
  'pullpush.sh',
  'conf',
  'deploy',
  'domain',
  'model',
  'assets/img',
]);

function findPages(dir = '.', sink = []) {
  if (ignorePath.has(dir.replace('./', '').replace('.\\', ''))) {
    debug && console.log('skip:', dir);
    return;
  }

  const files = readdirSync(dir).filter(f => f[0] !== '_');
  const svelteFiles = files.filter(f => f.endsWith('.svelte') && statSync(join(dir, f)).isFile());
  svelteFiles.forEach(f => sink.push(join(dir, f)));

  files
    .filter(f => !svelteFiles.includes(f))
    .map(f => join(dir, f))
    .filter(f => statSync(f).isDirectory())
    .forEach(f => findPages(f, sink));
  return sink;
}

const zIdPrefix = `z_placeholder_${Math.floor(Math.random() * 1000000000).toString(16)}_`;
const zReplacer = s => {
  const result = `"${zIdPrefix}${Buffer.from(s).toString('base64')}"`;
  debug && console.log('z-replace:', s, '->', result);
  return result;
};

const zPlaceholderReplacer = (content, isSSR = global.zIsSSR) => {
  if (!content) return content;
  if (isSSR) return content;
  return content
    .replace(/\#\{\s*\w+\s*\}/gs, zReplacer)
    .replace(/\/\*\!\s*\w+\s*\*\//gs, zReplacer)
    .replace(/\[\s*\/\*\s*\w+\s*\*\/\s*\]/gs, zReplacer)
    .replace(/\{\s*\/\*\s*\w+\s*\*\/\s*\}/gs, zReplacer);
};

global.zPlaceholderReplacer = zPlaceholderReplacer;

const zPlaceholderRestore = (content, sink) =>
  content?.replace(new RegExp(`("|')?${zIdPrefix}(\\w+=*)\\1?`, 'g'), (_, q, s) => {
    s = Buffer.from(s, 'base64').toString('ascii');
    sink.push(s);
    debug && console.log('z-restore', _, s);
    return s;
  });

const svelteJsPathResolver = {
  name: 'svelteJsPathResolver',
  setup(build) {
    const options = { filter: /\.svelte\.(ts)$/ };

    build.onResolve(options, ({ path, resolveDir }) => ({ path: join(resolveDir, path) }));
    build.onLoad(options, ({ path }) => ({
      contents: `
        import { mount } from 'svelte';
        import App from './${basename(path).replace(/\.ts$/, '')}';

        const target = document.getElementById('app');
        target.innerHTML = '';

        export const app = mount(App, { target });
        export default app;
      `,
      loader: 'ts',
      resolveDir: dirname(path),
    }));
  },
};

async function createBuilder(entryPoints) {
  console.log('pages:', entryPoints);

  const buildOptions = {
    entryPoints: entryPoints.map(s => `${s}.ts`),
    bundle: true,
    outdir: '.',
    write: false,
    plugins: [
      svelteJsPathResolver,
      sveltePlugin({ ...svelteConfig, compilerOptions: { ...svelteConfig.compilerOptions, css: 'injected' } }),
    ],
    sourcemap: false,
    minify,
  };

  if (watch) {
    const context = await esbuild.context(buildOptions);
    const buildResult = await context.rebuild();

    return {
      outputFiles: buildResult.outputFiles,
      rebuild: () => context.rebuild(),
      dispose: () => context.dispose(),
    };
  }

  return esbuild.build(buildOptions);
}

function layoutFor(path, content = {}) {
  path = (() => {
    let temp = join(path, '..', '_layout.html');

    while (true) {
      if (existsSync(temp)) return temp;
      if (resolve(__dirname) === resolve(dirname(temp))) return;

      temp = join(temp, '../..', '_layout.html');
    }
  })();

  layoutFor.cache = layoutFor.cache || {};

  const defaultKey = '_DEFAULT_LAYOUT';
  if (!path && layoutFor.cache[defaultKey]) return layoutFor.cache[defaultKey];

  const cache = layoutFor.cache[path];
  const m = statSync(path).mtimeMs;
  if (cache && m === cache.m) return cache;

  const tree = parse(
    path
      ? readFileSync(path, 'utf-8')
      : `<!DOCTYPE html>
<html>
  <head>
    <title>#{title}</title>
  </head>
  <body>
    <h1>layout not found, please create <b>_layout.html</b></h1>
    <slot></slot>
  </body>
</html>`
  );

  let slot = null;
  let body = null;
  let stack = [...tree.childNodes];

  while (stack.length && (slot == null || body == null)) {
    const t = stack.pop();
    if (t.nodeName === 'body') body = t;
    if (t.nodeName === 'slot' || (t.nodeName === '#comment' && t.data?.trim() === 'content_goes_here')) slot = t;
    t.childNodes && (stack = [...stack, ...t.childNodes]);
  }

  if (!body) throw new Error('body not found');

  const appKey = `${Math.random()}-APP-${Math.random()}`;
  if (slot) {
    slot.nodeName = 'main';
    slot.tagName = 'main';
    delete slot.data;
    slot.attrs = [{ name: 'id', value: 'app' }, ...(slot.attrs || []).filter(t => t.name !== 'id')];
    slot.childNodes = [{ nodeName: '#text', value: appKey }];
  } else {
    body.childNodes.push({
      nodeName: 'main',
      tagName: 'main',
      attrs: [{ name: 'id', value: 'app' }],
      childNodes: [{ nodeName: '#text', value: appKey }],
      namespaceURI: body.namespaceURI,
    });
  }

  const jsKey = `${Math.random()}-JS-${Math.random()}`;
  const cssKey = `${Math.random()}-CSS-${Math.random()}`;
  const comments = { nodeName: '#comment', data: '' };
  const cssVarsComments = { nodeName: '#comment', data: '' };
  const jsVarsComments = { nodeName: '#comment', data: '' };

  body.childNodes = [
    ...body.childNodes,
    comments,
    { nodeName: '#text', value: '\n' },
    logVars && cssVarsComments,
    { nodeName: '#text', value: '\n' },
    logVars && jsVarsComments,
    { nodeName: '#text', value: '\n' },
    {
      nodeName: 'style',
      tagName: 'style',
      attrs: [],
      childNodes: [{ nodeName: '#text', value: cssKey }],
      namespaceURI: body.namespaceURI,
    },
    { nodeName: '#text', value: '\n' },
    {
      nodeName: 'script',
      tagName: 'script',
      attrs: [],
      childNodes: [{ nodeName: '#text', value: jsKey }],
      namespaceURI: body.namespaceURI,
    },
    { nodeName: '#text', value: '\n' },
  ];

  debug && console.log('build layout for:', path || defaultKey);

  return (layoutFor.cache[path || defaultKey] = ({ js, css }) => {
    const cssVars = [];
    const jsVars = [];
    js = zPlaceholderRestore(js, jsVars) || '';
    css = zPlaceholderRestore(css, cssVars) || '';

    cssVarsComments.data = cssVars.length ? `--- CSS z-vars --- \n${cssVars.join('\n')}` : '';
    jsVarsComments.data = jsVars.length ? `--- JS z-vars --- \n${jsVars.join('\n')}` : '';

    const html = zPlaceholderRestore(content.html, []) || '';
    const innerCss = (content.css || {}).code || '';

    const result = serialize(tree)
      .replace(cssKey, css + innerCss)
      .replace(jsKey, js)
      .replace(appKey, html)
      .replaceAll(`fakecss:${__dirname}`, 'fakecss:.');

    if (debug && !result.includes(html)) {
      console.warn('WARNING: SSR HTML not found in final output for', path);
    }
    return result;
  });
}

(async () => {
  watch && console.log('first build start');
  const initialPages = findPages();
  global.zIsSSR = false;
  let builder = await createBuilder(initialPages);

  const compiledFiles = new Set();
  const cache = {};

  async function saveFiles(files = builder, layoutChanged = false) {
    const output = {};
    let unchanged = 0;

    const pages = Array.from(new Set(files.outputFiles.map(f => f.path.replace(/\.svelte\.\w+$/, '.svelte'))));
    global.zIsSSR = true;
    await esbuild.build({
      entryPoints: pages,
      bundle: true,
      platform: 'node',
      format: 'esm',
      outdir: '.ssr_temp',
      write: true,
      plugins: [sveltePlugin({ ...svelteConfig, compilerOptions: { ...svelteConfig.compilerOptions, generate: 'server' } })],
      external: ['svelte/server', 'axios'],
    });
    global.zIsSSR = false;

    for (const { path, text } of files.outputFiles) {
      const ext = /\.(\w+)$/.exec(path)?.[1];
      if (ext !== 'css' && ext !== 'js') throw new Error(`unknown ext:${ext}`);

      const key = path.replace(/\.svelte\.\w+$/, '');
      output[key] = output[key] || {};
      output[key][ext] = text;

      if (cache[path] === text && !layoutChanged) {
        unchanged += 1;
        continue;
      }
      cache[path] = text;
    }

    if (unchanged === files.outputFiles.length) return;
    if (Object.keys(output).length === 0) return console.log('no changes');

    for (const [path, data] of Object.entries(output)) {
      const ssrPath = resolve('.ssr_temp', relative(__dirname, path) + '.js');
      let html = '';
      try {
        const domNode = () => ({
          append() {},
          appendChild() {},
          addEventListener() {},
          blur() {},
          childNodes: [],
          classList: { add() {}, remove() {} },
          click() {},
          focus() {},
          getAttribute: () => null,
          innerHTML: '',
          parentNode: null,
          querySelector: () => null,
          remove() {},
          removeAttribute() {},
          removeChild() {},
          setAttribute() {},
          style: {},
        });

        globalThis.window ||= {
          location: { href: '', pathname: '', host: '' },
          navigator: { clipboard: { writeText() {} } },
        };
        globalThis.location ||= globalThis.window.location;
        globalThis.navigator ||= globalThis.window.navigator;
        globalThis.document ||= {
          body: domNode(),
          createElement: () => domNode(),
          getElementById: () => null,
          getElementsByTagName: () => [domNode()],
          querySelector: () => null,
          head: domNode(),
        };

        const { default: App } = await import(`${ssrPath}?${Date.now()}`);
        const rendered = render(App, { props: {} });
        html = rendered.html.replace(/<!--\[-->/g, '').replace(/<!--\]-->/g, '').replace(/<!---->/g, '');
      } catch (error) {
        console.error(`Failed to render SSR for ${path}:`, error);
      }

      const source = readFileSync(`${path}.svelte`, 'utf-8');
      const renderedSvelte = compile(source, { filename: `${path}.svelte`, css: 'external' });
      const content = layoutFor(path, { ...renderedSvelte, html })(data);

      const outPath = resolve(`${path}.html`);
      compiledFiles.add(outPath);
      console.log('compiled:', relative(resolve(__dirname), outPath));
      writeFileSync(outPath, content);
    }

    try {
      rmSync('.ssr_temp', { recursive: true, force: true });
    } catch {
      // ignore cleanup failure
    }
  }

  await saveFiles();
  watch && console.log('first build end');

  if (watch) {
    const pagePaths = new Set(initialPages.map(p => resolve(p)));
    let timeRef = null;

    function changeListener(path, stats, type, watcher) {
      switch (type) {
        case 'change':
          notifier.notify({
            title: 'Change occurs',
            message: `Change occurs in "${path}"`,
          });
          break;
        case 'add':
          notifier.notify({
            title: 'File added',
            message: `Added file "${path}"`,
          });
          break;
        case 'unlink':
          notifier.notify({
            title: 'File remove',
            message: `Removed file "${path}"`,
          });
          break;
      }

      if (compiledFiles.has(resolve(path))) return;
      console.log(`${type}:`, path.replace(__dirname, ''));

      const svelteFile = path[0] !== '_' && path.endsWith('.svelte');
      let pagesChanged = true;
      if (svelteFile && type === 'add') pagePaths.add(resolve(path));
      else if (svelteFile && type === 'unlink') pagePaths.delete(resolve(path));
      else pagesChanged = false;

      const layoutChanged = path.endsWith('_layout.html');

      if (timeRef) clearTimeout(timeRef);
      timeRef = setTimeout(async () => {
        if (pagesChanged) {
          if (builder.dispose) await builder.dispose();
          builder = await createBuilder(Array.from(pagePaths, p => relative(__dirname, p)));
          await saveFiles(builder, layoutChanged);
          return;
        }

        await saveFiles(await builder.rebuild(), layoutChanged);
      }, 200);
    }

    const watcher = chokidar
      .watch('.', { ignored: s => ignorePath.has(s) || ignorePath.has(join('./', s)), ignoreInitial: true })
      .on('change', (path, stats) => changeListener(path, stats, 'change', watcher))
      .on('add', (path, stats) => changeListener(path, stats, 'add', watcher))
      .on('unlink', (path, stats) => changeListener(path, stats, 'unlink', watcher))
      .on('ready', () => {
        console.log(`watching ${sum(Object.values(watcher.getWatched()).map(t => t.length))} files/dirs for changes`);
      })
      .on('error', err => console.log('ERROR:', err));
  }

  if (serve) {
    const FiveServerModule = await import('five-server');
    const FiveServer = FiveServerModule?.default?.default || FiveServerModule?.default || FiveServerModule;

    if (typeof FiveServer !== 'function') {
      throw new TypeError(`five-server export is not a constructor (typeof=${typeof FiveServer})`);
    }

    await new FiveServer().start({
      open: true,
      workspace: __dirname,
      ignore: [...ignorePath, /\.(js|ts|svelte)$/, /\_layout\.html$/],
      wait: 500,
    });
  }
})();

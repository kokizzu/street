import { writable } from 'svelte/store';
// translation related
import translation from '../translation.json';

export let isSideMenuOpen = writable(false); // Side Menu

let original = translation;
export let langOptions = {
    en: 'EN',
    tw: 'TW', // add more languages here
};

function translationStore() {
    let currentLang = '';
    if (window && window.localStorage) {
        currentLang = window.localStorage.getItem('lang') || '';
    }
    let kv = {currentLang: currentLang};
    let switchLang = () => {
        if (currentLang === 'EN' || currentLang === '') {
            for (let k in original) {
                kv[k] = original[k];
            }
        } else {
            for (let k in original) {
                if (original[k + currentLang])
                    kv[k] = original[k + currentLang];
            }
        }
    };
    const {subscribe, set, update} = writable(kv);

    return {
        currentLang,
        subscribe,
        changeLanguage: (newV) => {
            if (window && window.localStorage) {
                window.localStorage.setItem('lang', newV);
            }
            currentLang = newV;
            kv.currentLang = newV;
            switchLang();
            console.log('kv', kv);
            if (Object.keys(kv).length === 0) return;
            update((el) => kv);
        },

    };
}

export let T = translationStore();
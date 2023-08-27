<script>
    export let label = 'label';
    export let content = 'content';
</script>

<div class='pill-box pill-blue'>
    <span class='pill-box-left'>
      {label}
    </span>
    <span class='pill-box-right'>
      {content}
    </span>
</div>

<style>

    .pill-box {
        line-height : 25px;
    }

    .pill-box-left,
    .pill-box-right {
        border        : 1px solid #E6E6E6;
        border-radius : 4px;
        color         : #666;
        padding       : 2px;
    }

    .pill-box-left {
        background-color           : #E6E6E6;
        border-bottom-right-radius : 0;
        border-top-right-radius    : 0;
        padding-left               : 6px;
        text-align                 : center;
    }

    .pill-box-right {
        background-color          : #FFF;
        border-bottom-left-radius : 0;
        border-top-left-radius    : 0;
        border-left               : 0;
        padding-left              : 4px;
        padding-right             : 6px;
    }

    .pill-blue .pill-box-left,
    .pill-blue .pill-box-right {
        border-color : #4DA6FF;
    }

    .pill-blue .pill-box-left {
        background-color : #4DA6FF;
        color            : #E5F2FF;
    }
</style>
declare const _default: {
    '.uk-modal': {
        display: string;
        position: string;
        top: string;
        right: string;
        bottom: string;
        left: string;
        zIndex: string;
        overflowY: string;
        backgroundColor: string;
        opacity: string;
        transition: string;
        '@apply backdrop-blur-sm p-4': {};
    };
    '.uk-modal.uk-open': {
        opacity: string;
    };
    '.uk-modal-page': {
        overflow: string;
    };
    '.uk-modal-dialog': {
        position: string;
        boxSizing: string;
        margin: string;
        width: string;
        maxWidth: string;
        opacity: string;
        transform: string;
        transition: string;
        transitionProperty: string;
        '@apply overflow-hidden border border-border bg-background rounded-lg': {};
    };
    '.uk-open > .uk-modal-dialog': {
        opacity: string;
        transform: string;
    };
    '.uk-modal-container .uk-modal-dialog': {
        width: string;
    };
    '.uk-modal-full': {
        padding: string;
        background: string;
    };
    '.uk-modal-full .uk-modal-dialog': {
        margin: string;
        width: string;
        maxWidth: string;
        transform: string;
        '@apply w-full rounded-none border-none': {};
    };
    '.uk-modal-body': {
        display: string;
        '@apply p-6': {};
    };
    '.uk-modal-header': {
        display: string;
        '@apply p-6': {};
    };
    '.uk-modal-footer': {
        display: string;
        '@apply p-6': {};
    };
    '.uk-modal-body > :last-child, .uk-modal-header > :last-child, .uk-modal-footer > :last-child': {
        marginBottom: string;
    };
    '.uk-modal-title': {
        '@apply text-lg font-semibold leading-none tracking-tight': {};
    };
    "[class*='uk-modal-close-']": {
        position: string;
        zIndex: string;
        top: string;
        right: string;
        '@apply top-4 right-4': {};
    };
    "[class*='uk-modal-close-']:first-child + *": {
        marginTop: string;
    };
    '.uk-modal-close-outside': {
        top: string;
        right: string;
        transform: string;
        color: string;
    };
    '.uk-modal-close-outside:hover': {
        color: string;
    };
    '.uk-modal-close-full': {
        top: string;
        right: string;
    };
    '.uk-modal-header + .uk-modal-body, .uk-modal-body + .uk-modal-footer': {
        '@apply pt-0': {};
    };
    '.uk-modal-header ~ .uk-modal-footer': {
        '@apply pt-0': {};
    };
    '.uk-modal-header + .uk-modal-body.uk-overflow-auto': {
        '@apply pb-0': {};
    };
    '.uk-modal-body.uk-overflow-auto + .uk-modal-footer': {
        '@apply pt-6': {};
    };
    '.uk-flex': {
        display: string;
    };
    '.uk-flex-top': {
        '@apply items-start': {};
    };
    '.uk-margin-auto-vertical': {
        marginTop: string;
        marginBottom: string;
    };
    '.uk-overflow-auto': {
        overflow: string;
    };
    '.uk-overflow-auto > :last-child': {
        marginBottom: string;
    };
};
export default _default;

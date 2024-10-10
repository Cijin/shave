declare const _default: {
    '.uk-offcanvas': {
        display: string;
        position: string;
        top: string;
        bottom: string;
        left: string;
        zIndex: string;
    };
    '.uk-offcanvas-flip .uk-offcanvas': {
        right: string;
        left: string;
    };
    '.uk-offcanvas-bar': {
        position: string;
        top: string;
        bottom: string;
        left: string;
        boxSizing: string;
        width: string;
        overflowY: string;
        '@apply border-r border-border bg-background': {};
    };
    '.uk-offcanvas-flip .uk-offcanvas-bar': {
        left: string;
        right: string;
        borderRight: string;
        '@apply border-l border-border': {};
    };
    '.uk-open > .uk-offcanvas-bar': {
        left: string;
    };
    '.uk-offcanvas-flip .uk-open > .uk-offcanvas-bar': {
        left: string;
        right: string;
    };
    '.uk-offcanvas-bar-animation': {
        transition: string;
    };
    '.uk-offcanvas-flip .uk-offcanvas-bar-animation': {
        transitionProperty: string;
    };
    '.uk-offcanvas-reveal': {
        position: string;
        top: string;
        bottom: string;
        left: string;
        width: string;
        overflow: string;
        transition: string;
    };
    '.uk-offcanvas-reveal .uk-offcanvas-bar': {
        left: string;
    };
    '.uk-offcanvas-flip .uk-offcanvas-reveal .uk-offcanvas-bar': {
        left: string;
        right: string;
    };
    '.uk-open > .uk-offcanvas-reveal': {
        width: string;
    };
    '.uk-offcanvas-flip .uk-offcanvas-reveal': {
        right: string;
        left: string;
    };
    '.uk-offcanvas-close': {
        position: string;
        zIndex: string;
        top: string;
        right: string;
        padding: string;
    };
    '.uk-offcanvas-close:first-child + *': {
        marginTop: string;
    };
    '.uk-offcanvas-overlay': {
        width: string;
        touchAction: string;
    };
    '.uk-offcanvas-overlay::before': {
        content: string;
        position: string;
        top: string;
        bottom: string;
        left: string;
        right: string;
        backgroundColor: string;
        opacity: string;
        transition: string;
        '@apply backdrop-blur-sm': {};
    };
    '.uk-offcanvas-overlay.uk-open::before': {
        opacity: string;
    };
    '.uk-offcanvas-page, .uk-offcanvas-container': {
        overflowX: string[];
    };
    '.uk-offcanvas-container': {
        position: string;
        left: string;
        transition: string;
        boxSizing: string;
        width: string;
    };
    ':not(.uk-offcanvas-flip).uk-offcanvas-container-animation': {
        left: string;
    };
    '.uk-offcanvas-flip.uk-offcanvas-container-animation': {
        left: string;
    };
};
export default _default;

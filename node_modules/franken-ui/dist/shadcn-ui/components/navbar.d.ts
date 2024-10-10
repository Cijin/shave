declare const _default: {
    '.uk-navbar': {
        display: string;
        position: string;
        '--uk-navbar-dropbar-behind-color': string;
    };
    '.uk-navbar-container:not(.uk-navbar-transparent)': {
        '@apply border-b border-border/40 bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60': {};
    };
    ".uk-navbar-left, .uk-navbar-right, [class*='uk-navbar-center']": {
        display: string;
        alignItems: string;
        '@apply gap-6': {};
    };
    '.uk-navbar-right': {
        marginLeft: string;
    };
    '.uk-navbar-center:only-child': {
        marginLeft: string;
        marginRight: string;
        position: string;
    };
    '.uk-navbar-center:not(:only-child)': {
        position: string;
        top: string;
        left: string;
        transform: string;
        width: string;
        boxSizing: string;
        zIndex: string;
    };
    '.uk-navbar-center-left, .uk-navbar-center-right': {
        position: string;
        top: string;
    };
    '.uk-navbar-center-left': {
        right: string;
    };
    '.uk-navbar-center-right': {
        left: string;
    };
    "[class*='uk-navbar-center-']": {
        width: string;
        boxSizing: string;
    };
    '.uk-navbar-nav': {
        display: string;
        margin: string;
        padding: string;
        listStyle: string;
        '@apply gap-6': {};
    };
    '.uk-navbar-left, .uk-navbar-right, .uk-navbar-center:only-child': {
        flexWrap: string;
    };
    '.uk-navbar-nav > li > a, .uk-navbar-item, .uk-navbar-toggle': {
        display: string;
        justifyContent: string;
        alignItems: string;
        columnGap: string;
        boxSizing: string;
        fontSize: string;
        textDecoration: string;
        '@apply min-h-14': {};
    };
    '.uk-navbar-nav > li > a': {
        padding: string;
        transition: string;
        transitionProperty: string;
        '@apply text-foreground/60': {};
    };
    '.uk-navbar-nav > li:hover > a, .uk-navbar-nav > li > a[aria-expanded="true"]': {
        '@apply text-foreground/80': {};
    };
    '.uk-navbar-nav > li.uk-active > a': {
        '@apply text-foreground': {};
    };
    '.uk-navbar-parent-icon': {
        marginLeft: string;
        transition: string;
    };
    '.uk-navbar-nav > li > a[aria-expanded="true"] .uk-navbar-parent-icon': {
        transform: string;
    };
    '.uk-navbar-item': {
        padding: string;
        '@apply text-foreground/60': {};
    };
    '.uk-navbar-item > :last-child': {
        marginBottom: string;
    };
    '.uk-navbar-toggle': {
        padding: string;
        '@apply text-foreground/60': {};
    };
    '.uk-navbar-toggle:hover, .uk-navbar-toggle[aria-expanded="true"]': {
        textDecoration: string;
        '@apply text-foreground/80': {};
    };
    '.uk-navbar-subtitle': {
        fontSize: string;
    };
    '.uk-navbar-justify .uk-navbar-left, .uk-navbar-justify .uk-navbar-right, .uk-navbar-justify .uk-navbar-nav, .uk-navbar-justify .uk-navbar-nav > li, .uk-navbar-justify .uk-navbar-item, .uk-navbar-justify .uk-navbar-toggle': {
        flexGrow: string;
    };
    '.uk-navbar-dropdown': {
        '--uk-position-offset': string;
        '--uk-position-shift-offset': string;
        '--uk-position-viewport-offset': string;
        width: string;
        '@apply rounded-md border border-border bg-popover text-popover-foreground shadow-md': {};
    };
    '.uk-navbar-dropdown > :last-child': {
        marginBottom: string;
    };
    '.uk-navbar-dropdown :focus-visible': {
        outlineColor: string;
    };
    '.uk-navbar-dropdown .uk-drop-grid': {
        marginLeft: string;
    };
    '.uk-navbar-dropdown .uk-drop-grid > *': {
        paddingLeft: string;
    };
    '.uk-navbar-dropdown .uk-drop-grid > .uk-grid-margin': {
        marginTop: string;
    };
    '.uk-navbar-dropdown-width-2:not(.uk-drop-stack)': {
        width: string;
    };
    '.uk-navbar-dropdown-width-3:not(.uk-drop-stack)': {
        width: string;
    };
    '.uk-navbar-dropdown-width-4:not(.uk-drop-stack)': {
        width: string;
    };
    '.uk-navbar-dropdown-width-5:not(.uk-drop-stack)': {
        width: string;
    };
    '.uk-navbar-dropdown-dropbar': {
        width: string;
        background: string;
        padding: string;
        '--uk-position-offset': string;
        '--uk-position-shift-offset': string;
        '--uk-position-viewport-offset': string;
        boxShadow: string;
    };
    '.uk-navbar-dropdown-nav': {
        fontSize: string;
    };
    '.uk-navbar-dropdown-nav > li > a': {
        '@apply m-1 flex select-none items-center rounded-sm px-2 py-1.5 text-sm': {};
    };
    '.uk-navbar-dropdown-nav > li > a:hover': {
        '@apply bg-accent text-accent-foreground': {};
    };
    '.uk-navbar-dropdown-nav > li.uk-active > a': {
        '@apply bg-secondary text-secondary-foreground': {};
    };
    '.uk-navbar-dropdown-nav .uk-nav-subtitle': {
        fontSize: string;
    };
    '.uk-navbar-dropdown-nav .uk-nav-header': {
        '@apply m-1 px-2 py-1.5': {};
    };
    '.uk-navbar-dropdown-nav .uk-nav-divider': {
        borderTop: string;
        '@apply border-border': {};
    };
    '.uk-navbar-dropdown-nav .uk-nav-sub a': {
        '@apply ml-1 px-2 py-1.5': {};
    };
    '.uk-navbar-dropdown-nav .uk-nav-sub a:hover': {
        '@apply bg-accent text-accent-foreground': {};
    };
    '.uk-navbar-dropdown-nav .uk-nav-sub li.uk-active > a': {
        '@apply bg-secondary text-secondary-foreground': {};
    };
    '.uk-navbar-container': {
        transition: string;
        transitionProperty: string;
    };
    '.uk-navbar-toggle-icon': {
        '@apply text-foreground/60': {};
    };
    ':hover > .uk-navbar-toggle-icon': {
        '@apply text-foreground/80': {};
    };
};
export default _default;

declare const _default: {
    '.uk-tab': {
        display: string;
        flexWrap: string;
        padding: string;
        listStyle: string;
        position: string;
        '@apply text-muted-foreground': {};
    };
    '.uk-tab::before': {
        content: string;
        position: string;
        bottom: string;
        left: string;
        right: string;
        borderBottom: string;
        '@apply border-border': {};
    };
    '.uk-tab > *': {
        flex: string;
        position: string;
    };
    '.uk-tab > * > a': {
        display: string;
        alignItems: string;
        columnGap: string;
        justifyContent: string;
        padding: string;
        borderBottom: string;
        fontSize: string;
        textTransform: string;
        transition: string;
        '@apply font-semibold text-muted-foreground': {};
    };
    '.uk-tab > * > a:hover': {
        textDecoration: string;
    };
    '.uk-tab > .uk-active > a': {
        '@apply border-primary bg-background text-foreground': {};
    };
    '.uk-tab > .uk-disabled > a': {
        '@apply opacity-50 disabled:pointer-events-none': {};
    };
    '.uk-tab-bottom::before': {
        top: string;
        bottom: string;
    };
    '.uk-tab-bottom > * > a': {
        borderTop: string;
        borderBottom: string;
    };
    '.uk-tab-left, .uk-tab-right': {
        flexDirection: string;
        marginLeft: string;
    };
    '.uk-tab-left > *, .uk-tab-right > *': {
        paddingLeft: string;
    };
    '.uk-tab-left::before': {
        top: string;
        bottom: string;
        left: string;
        right: string;
        borderLeft: string;
        borderBottom: string;
        '@apply border-border': {};
    };
    '.uk-tab-right::before': {
        top: string;
        bottom: string;
        left: string;
        right: string;
        borderLeft: string;
        borderBottom: string;
        '@apply border-border': {};
    };
    '.uk-tab-left > * > a': {
        justifyContent: string;
        borderRight: string;
        borderBottom: string;
    };
    '.uk-tab-right > * > a': {
        justifyContent: string;
        borderLeft: string;
        borderBottom: string;
    };
    '.uk-tab .uk-dropdown': {
        marginLeft: string;
    };
    '.uk-tab-alt': {
        '@apply flex h-9 w-full items-center justify-center rounded-lg bg-muted p-1 text-muted-foreground': {};
    };
    '.uk-tab-alt > *': {
        '@apply flex-grow': {};
    };
    '.uk-tab-alt > * > a': {
        '@apply inline-flex h-auto w-full items-center justify-center whitespace-nowrap px-3 py-1 text-sm font-medium no-underline': {};
    };
    '.uk-tab-alt > .uk-active > a': {
        '@apply rounded-md border-b-0 bg-background text-foreground': {};
    };
};
export default _default;

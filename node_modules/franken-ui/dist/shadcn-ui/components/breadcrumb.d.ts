declare const _default: {
    '.uk-breadcrumb': {
        padding: string;
        listStyle: string;
        '@apply text-sm': {};
    };
    '.uk-breadcrumb > *': {
        display: string;
    };
    '.uk-breadcrumb > * > *': {
        '@apply inline-flex items-center text-muted-foreground': {};
    };
    '.uk-breadcrumb > * > :hover': {
        textDecoration: string;
        '@apply text-foreground': {};
    };
    '.uk-breadcrumb > :last-child > span, .uk-breadcrumb > :last-child > a:not([href])': {
        '@apply font-medium text-foreground': {};
    };
    '.uk-breadcrumb > :nth-child(n + 2):not(.uk-first-column)::before': {
        content: string;
        display: string;
        fontSize: string;
        '@apply ml-2 mr-3 text-muted-foreground': {};
    };
    '.uk-breadcrumb > .uk-disabled > *': {
        '@apply opacity-50': {};
    };
};
export default _default;

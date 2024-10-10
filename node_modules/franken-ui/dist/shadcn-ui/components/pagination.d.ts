declare const _default: {
    '.uk-pagination': {
        display: string;
        flexWrap: string;
        alignItems: string;
        listStyle: string;
        '@apply gap-1': {};
    };
    '.uk-pagination > *': {
        flex: string;
        position: string;
    };
    '.uk-pagination > * > *': {
        columnGap: string;
        transition: string;
        '@apply inline-flex h-9 min-w-9 items-center justify-center whitespace-nowrap rounded-md': {};
    };
    '.uk-pagination > * > :hover': {
        textDecoration: string;
        '@apply bg-accent text-accent-foreground': {};
    };
    '.uk-pagination > .uk-active > *': {
        '@apply border border-input shadow': {};
    };
    '.uk-pagination > .uk-disabled > *': {
        '@apply opacity-50': {};
    };
};
export default _default;

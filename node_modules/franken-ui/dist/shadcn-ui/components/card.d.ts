declare const _default: {
    '.uk-card': {
        position: string;
        boxSizing: string;
        '@apply rounded-xl border border-border shadow': {};
    };
    '.uk-card-body': {
        display: string;
        '@apply p-6': {};
    };
    '.uk-card-header': {
        display: string;
        '@apply p-6': {};
    };
    '.uk-card-footer': {
        display: string;
        '@apply p-6': {};
    };
    '.uk-card-body > :last-child, .uk-card-header > :last-child, .uk-card-footer > :last-child': {
        marginBottom: string;
    };
    '.uk-card-title': {
        '@apply font-semibold leading-none tracking-tight': {};
    };
    '.uk-card-default': {
        '@apply bg-card text-card-foreground': {};
    };
    '.uk-card-default .uk-card-header': {
        borderBottom: string;
        '@apply border-border': {};
    };
    '.uk-card-default .uk-card-footer': {
        borderTop: string;
        '@apply border-border': {};
    };
    '.uk-card-primary': {
        '@apply bg-primary text-primary-foreground': {};
    };
    '.uk-card-secondary': {
        '@apply bg-secondary text-secondary-foreground': {};
    };
    '.uk-card-danger': {
        '@apply bg-destructive text-destructive-foreground': {};
    };
    '.uk-card-header + .uk-card-body, .uk-card-body + .uk-card-footer': {
        '@apply pt-0': {};
    };
    '.uk-card-header ~ .uk-card-footer': {
        '@apply pt-0': {};
    };
};
export default _default;

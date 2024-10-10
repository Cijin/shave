declare const _default: {
    '.uk-notification': {
        position: string;
        top: string;
        left: string;
        zIndex: string;
        boxSizing: string;
        width: string;
    };
    '.uk-notification-top-right, .uk-notification-bottom-right': {
        left: string;
        right: string;
    };
    '.uk-notification-top-center, .uk-notification-bottom-center': {
        left: string;
        marginLeft: string;
    };
    '.uk-notification-bottom-left, .uk-notification-bottom-right, .uk-notification-bottom-center': {
        top: string;
        bottom: string;
    };
    '.uk-notification-message': {
        position: string;
        cursor: string;
        '@apply rounded-md border border-border bg-background p-4 pr-6 text-sm text-foreground shadow-lg': {};
    };
    '* + .uk-notification-message': {
        marginTop: string;
    };
    '.uk-notification-close': {
        display: string;
        position: string;
        '@apply right-1 top-1 p-1 text-foreground/50': {};
    };
    '.uk-notification-message:hover .uk-notification-close': {
        display: string;
    };
    '.uk-notification-message-danger': {
        '@apply border-destructive bg-destructive text-destructive-foreground': {};
    };
    '.uk-notification-message-danger .uk-notification-close': {
        '@apply text-destructive-foreground/50': {};
    };
};
export default _default;

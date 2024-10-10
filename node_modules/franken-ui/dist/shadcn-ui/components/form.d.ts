declare const _default: {
    '.uk-input, .uk-select, .uk-textarea, .uk-radio, .uk-checkbox': {
        boxSizing: string;
        margin: string;
        borderRadius: string;
        font: string;
    };
    '.uk-input': {
        overflow: string;
    };
    '.uk-select': {
        textTransform: string;
    };
    '.uk-select optgroup': {
        font: string;
        fontWeight: string;
    };
    '.uk-textarea': {
        overflow: string;
    };
    '.uk-input[type="search"]::-webkit-search-cancel-button, .uk-input[type="search"]::-webkit-search-decoration': {
        WebkitAppearance: string;
    };
    '.uk-input[type="number"]::-webkit-inner-spin-button, .uk-input[type="number"]::-webkit-outer-spin-button': {
        height: string;
    };
    '.uk-input::-moz-placeholder, .uk-textarea::-moz-placeholder': {
        opacity: string;
    };
    '.uk-radio:not(:disabled), .uk-checkbox:not(:disabled)': {
        cursor: string;
    };
    '.uk-fieldset': {
        border: string;
        margin: string;
        padding: string;
        minWidth: string;
    };
    '.uk-input, .uk-textarea': {
        WebkitAppearance: string;
    };
    '.uk-input, .uk-select, .uk-textarea': {
        maxWidth: string;
        width: string;
        '@apply rounded-md border border-input bg-transparent px-3 py-1 text-sm text-foreground shadow-sm': {};
    };
    '.uk-input, .uk-select:not([multiple]):not([size])': {
        verticalAlign: string;
        display: string;
        '@apply h-9': {};
    };
    '.uk-input:not(input), .uk-select:not(select)': {
        lineHeight: string;
    };
    '.uk-select[multiple], .uk-select[size], .uk-textarea': {
        paddingTop: string;
        paddingBottom: string;
        verticalAlign: string;
        '@apply min-h-[60px] py-2': {};
    };
    '.uk-select[multiple], .uk-select[size]': {
        resize: string;
    };
    '.uk-input:focus, .uk-select:focus, .uk-textarea:focus': {
        outline: string;
        '@apply outline-none ring-1 ring-ring': {};
    };
    '.uk-input:disabled, .uk-select:disabled, .uk-textarea:disabled': {
        '@apply opacity-50': {};
    };
    '.uk-input::placeholder': {
        '@apply text-muted-foreground': {};
    };
    '.uk-textarea::placeholder': {
        '@apply text-muted-foreground': {};
    };
    '.uk-form-small': {
        '@apply text-xs': {};
    };
    '.uk-form-small:not(textarea):not([multiple]):not([size])': {
        '@apply h-8 px-3 py-0': {};
    };
    '.uk-form-danger, .uk-form-danger:focus': {
        '@apply ring-destructive': {};
    };
    '.uk-form-blank': {
        background: string;
        borderColor: string;
        '@apply shadow-none': {};
    };
    '.uk-form-blank:focus': {
        '@apply ring-0': {};
    };
    '.uk-select:not([multiple]):not([size])': {
        WebkitAppearance: string;
        MozAppearance: string;
        paddingRight: string;
        backgroundImage: string;
        backgroundRepeat: string;
        backgroundPosition: string;
    };
    '.uk-select:not([multiple]):not([size]) option': {
        color: string;
    };
    '.uk-input[list]': {
        paddingRight: string;
        backgroundRepeat: string;
        backgroundPosition: string;
    };
    '.uk-input[list]:hover, .uk-input[list]:focus': {
        backgroundImage: string;
    };
    '.uk-input[list]::-webkit-calendar-picker-indicator': {
        display: string;
    };
    '.uk-radio, .uk-checkbox': {
        display: string;
        height: string;
        width: string;
        overflow: string;
        marginTop: string;
        verticalAlign: string;
        WebkitAppearance: string;
        MozAppearance: string;
        backgroundRepeat: string;
        backgroundPosition: string;
        '@apply rounded-sm border border-primary shadow': {};
    };
    '.uk-radio': {
        borderRadius: string;
    };
    '.uk-radio:focus, .uk-checkbox:focus': {
        outline: string;
        '@apply outline-none ring-1 ring-ring': {};
    };
    '.uk-radio:checked:focus, .uk-checkbox:checked:focus, .uk-checkbox:indeterminate:focus': {
        '@apply outline-none ring-1 ring-ring': {};
    };
    '.uk-radio:checked': {
        backgroundImage: string;
    };
    '.uk-checkbox:checked': {
        backgroundImage: string;
    };
    '.uk-checkbox:indeterminate': {
        backgroundImage: string;
    };
    '.uk-form-custom': {
        display: string;
        position: string;
        maxWidth: string;
        verticalAlign: string;
    };
    '.uk-form-custom select, .uk-form-custom input[type="file"]': {
        position: string;
        top: string;
        zIndex: string;
        width: string;
        height: string;
        left: string;
        WebkitAppearance: string;
        opacity: string;
        cursor: string;
    };
    '.uk-form-custom input[type="file"]': {
        fontSize: string;
        overflow: string;
    };
    '.uk-form-label': {
        fontSize: string;
        '@apply text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70': {};
    };
    '.uk-form-stacked .uk-form-label': {
        display: string;
        marginBottom: string;
    };
    '.uk-form-icon': {
        position: string;
        top: string;
        bottom: string;
        left: string;
        width: string;
        display: string;
        justifyContent: string;
        alignItems: string;
    };
    '.uk-form-icon:not(a):not(button):not(input)': {
        pointerEvents: string;
    };
    '.uk-form-icon:not(.uk-form-icon-flip) ~ .uk-input': {
        paddingLeft: string;
    };
    '.uk-form-icon-flip': {
        right: string;
        left: string;
    };
    '.uk-form-icon-flip ~ .uk-input': {
        paddingRight: string;
    };
    '.uk-toggle-switch': {
        WebkitAppearance: string;
        MozAppearance: string;
        appearance: string;
        position: string;
        boxSizing: string;
        borderRadius: string;
        '@apply h-5 w-9 bg-input shadow-sm': {};
    };
    '.uk-toggle-switch::before': {
        content: string;
        position: string;
        top: string;
        left: string;
        transform: string;
        boxSizing: string;
        borderRadius: string;
        transitionProperty: string;
        transitionDuration: string;
        transitionTimingFunction: string;
        marginLeft: string;
        '@apply h-4 w-4 bg-background': {};
    };
    '.uk-toggle-switch:checked::before': {
        '@apply left-4': {};
    };
    '.uk-toggle-switch:disabled': {
        opacity: string;
    };
    '.uk-toggle-switch-primary:checked': {
        '@apply bg-primary': {};
    };
    '.uk-toggle-switch-danger:checked': {
        '@apply bg-destructive': {};
    };
    '.uk-form-help': {
        fontSize: string;
        '@apply font-medium': {};
    };
    "[class*='uk-inline']": {
        display: string;
        position: string;
        maxWidth: string;
        verticalAlign: string;
        WebkitBackfaceVisibility: string;
    };
};
export default _default;

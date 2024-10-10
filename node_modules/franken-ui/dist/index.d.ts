export type Colors = {
    '--background': string;
    '--foreground': string;
    '--muted': string;
    '--muted-foreground': string;
    '--popover': string;
    '--popover-foreground': string;
    '--card': string;
    '--card-foreground': string;
    '--border': string;
    '--input': string;
    '--primary': string;
    '--primary-foreground': string;
    '--secondary': string;
    '--secondary-foreground': string;
    '--accent': string;
    '--accent-foreground': string;
    '--destructive': string;
    '--destructive-foreground': string;
    '--ring': string;
};
export type Options = {
    customPalette?: {
        [key: string]: Colors;
    };
};
declare const _default: {
    (options: Options): {
        handler: import("tailwindcss/types/config.js").PluginCreator;
        config?: Partial<import("tailwindcss/types/config.js").Config> | undefined;
    };
    __isOptionsFunction: true;
};
export default _default;

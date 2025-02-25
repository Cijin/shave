import franken from 'franken-ui/shadcn-ui/preset-quick';

/** @type {import('tailwindcss').Config} */
export default {
	  presets: [franken({
      customPalette: {
        ".uk-theme-blue":{
          "--background": "169 44% 99%",
          "--foreground": "169 63% 3%",
          "--muted": "169 27% 91%",
          "--muted-foreground": "169 3% 36%",
          "--popover": "169 44% 98%",
          "--popover-foreground": "169 63% 2%",
          "--card": "169 44% 98%",
          "--card-foreground": "169 63% 2%",
          "--border": "169 13% 90%",
          "--input": "169 13% 90%",
          "--primary": "169 85% 63%",
          "--primary-foreground": "169 85% 3%",
          "--secondary": "169 8% 81%",
          "--secondary-foreground": "169 8% 21%",
          "--accent": "169 8% 81%",
          "--accent-foreground": "169 8% 21%",
          "--destructive": "12 98% 39%",
          "--destructive-foreground": "0 0% 100%",
          "--ring": "169 85% 63%",
          "--chart-1": "169 85% 63%",
          "--chart-2": "169 8% 81%",
          "--chart-3": "169 8% 81%",
          "--chart-4": "169 8% 84%",
          "--chart-5": "169 88% 63%",
        },
        ".dark.uk-theme.blue": {
          "--background": "169 44% 1%",
          "--foreground": "169 17% 99%",
          "--muted": "169 27% 9%",
          "--muted-foreground": "169 3% 64%",
          "--popover": "169 44% 2%",
          "--popover-foreground": "0 0% 100%",
          "--card": "169 44% 2%",
          "--card-foreground": "0 0% 100%",
          "--border": "169 13% 10%",
          "--input": "169 13% 10%",
          "--primary": "169 85% 63%",
          "--primary-foreground": "169 85% 3%",
          "--secondary": "169 2% 12%",
          "--secondary-foreground": "169 2% 72%",
          "--accent": "169 2% 12%",
          "--accent-foreground": "169 2% 72%",
          "--destructive": "12 98% 56%",
          "--destructive-foreground": "0 0% 100%",
          "--ring": "169 85% 63%",
          "--chart-1": "169 85% 63%",
          "--chart-2": "169 2% 12%",
          "--chart-3": "169 2% 12%",
          "--chart-4": "169 2% 15%",
          "--chart-5": "169 88% 63%"
        },
    },
  })],
  content: ["./views/**/*.html", "./views/**/*.templ", "./views/**/*.go",],
	safelist: [
		{
			pattern: /^uk-/
		},
		'ProseMirror',
		'ProseMirror-focused',
		'tiptap'
	],
	theme: {
		extend: {}
	},
	plugins: []
};

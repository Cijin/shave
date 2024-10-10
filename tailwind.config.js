import franken from 'franken-ui/shadcn-ui/preset-quick';

/** @type {import('tailwindcss').Config} */
export default {
	  presets: [franken({
    customPalette: {
      ".uk-theme-blue": {
        "--background": "0.00 0.00% 100.00%",
        "--foreground": "216.00 100.00% 4.90%",
        "--primary": "216.47 100.00% 50.00%",
        "--success": "142 76% 45%",
        "--primary-foreground": "216.00 100.00% 95.10%",
        "--card": "216.92 100.00% 97.45%",
        "--card-foreground": "216.00 100.00% 4.90%",
        "--popover": "0.00 0.00% 100.00%",
        "--popover-foreground": "216.00 100.00% 4.90%",
        "--secondary": "216.32 100.00% 92.55%",
        "--secondary-foreground": "0.00 0.00% 0.00%",
        "--muted": "216.00 100.00% 95.10%",
        "--muted-foreground": "0.00 0.00% 40.00%",
        "--accent": "216.00 100.00% 95.10%",
        "--accent-foreground": "216.32 100.00% 7.45%",
        "--destructive": "0 84.2% 60.2%",
        "--destructive-foreground": "210 40% 98%",
        "--border": "216.47 100.00% 90.00%",
        "--input": "215.63 33.33% 81.18%",
        "--ring": "216.47 100.00% 50.00%",
        "--chart-1": "216.47 100.00% 50.00%",
        "--chart-2": "216.50 100.00% 57.45%",
        "--chart-3": "216.51 100.00% 67.45%",
        "--chart-4": "216.43 100.00% 72.55%",
        "--chart-5": "216.52 100.00% 77.45%",
      },
      ".dark.uk-theme.blue": {
        "--background": "0.00 0.00% 0.00%",
        "--foreground": "216.32 100.00% 92.55%",
        "--primary": "216.47 100.00% 50.00%",
        "--success": "142 76% 28%",
        "--primary-foreground": "216.00 100.00% 95.10%",
        "--card": "216.00 33.33% 5.88%",
        "--card-foreground": "216.32 100.00% 92.55%",
        "--popover": "0.00 0.00% 0.00%",
        "--popover-foreground": "216.32 100.00% 92.55%",
        "--secondary": "216.40 100.00% 17.45%",
        "--secondary-foreground": "0.00 0.00% 100.00%",
        "--muted": "0.00 0.00% 9.80%",
        "--muted-foreground": "0.00 0.00% 50.20%",
        "--accent": "216.40 100.00% 17.45%",
        "--accent-foreground": "216.32 100.00% 92.55%",
        "--destructive": "0 84.2% 60.2%",
        "--destructive-foreground": "210 40% 98%",
        "--border": "216.47 100.00% 10.00%",
        "--input": "216.92 25.49% 20.00%",
        "--ring": "216.47 100.00% 50.00%",
        "--chart-1": "216.47 100.00% 50.00%",
        "--chart-2": "216.50 100.00% 42.55%",
        "--chart-3": "216.51 100.00% 32.55%",
        "--chart-4": "216.43 100.00% 27.45%",
        "--chart-5": "216.52 100.00% 22.55%",
      }
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

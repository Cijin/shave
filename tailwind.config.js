import franken from 'franken-ui/shadcn-ui/preset-quick';

/** @type {import('tailwindcss').Config} */
export default {
	  presets: [franken({
    customPalette: {
      ".uk-theme-green": {
        "--background": "0.00 0.00% 100.00%",
        "--foreground": "144.55 84.62% 5.10%",
        "--primary": "144.00 91.09% 51.57%",
        "--primary-foreground": "144.55 84.62% 5.10%",
        "--card": "147.27 84.62% 97.45%",
        "--card-foreground": "144.55 84.62% 5.10%",
        "--popover": "0.00 0.00% 100.00%",
        "--popover-foreground": "144.55 84.62% 5.10%",
        "--secondary": "143.64 89.19% 92.75%",
        "--secondary-foreground": "0.00 0.00% 0.00%",
        "--muted": "143.48 92.00% 95.10%",
        "--muted-foreground": "0.00 0.00% 40.00%",
        "--accent": "143.48 92.00% 95.10%",
        "--accent-foreground": "142.94 85.00% 7.84%",
        "--destructive": "0 84.2% 60.2%",
        "--destructive-foreground": "210 40% 98%",
        "--border": "144.00 91.84% 90.39%",
        "--input": "143.57 29.79% 81.57%",
        "--ring": "144.00 91.09% 51.57%",
        "--chart-1": "144.00 91.09% 51.57%",
        "--chart-2": "144.06 91.43% 58.82%",
        "--chart-3": "143.84 91.25% 68.63%",
        "--chart-4": "144.19 91.18% 73.33%",
        "--chart-5": "143.76 90.99% 78.24%"
      },
      ".dark.uk-theme.green": {
        "--background": "0.00 0.00% 0.00%",
        "--foreground": "143.64 89.19% 92.75%",
        "--primary": "144.00 91.09% 51.57%",
        "--primary-foreground": "144.55 84.62% 5.10%",
        "--card": "146.67 31.03% 5.69%",
        "--card-foreground": "143.64 89.19% 92.75%",
        "--popover": "0.00 0.00% 0.00%",
        "--popover-foreground": "143.64 89.19% 92.75%",
        "--secondary": "143.85 84.78% 18.04%",
        "--secondary-foreground": "0.00 0.00% 100.00%",
        "--muted": "0.00 0.00% 9.80%",
        "--muted-foreground": "0.00 0.00% 50.20%",
        "--accent": "143.85 84.78% 18.04%",
        "--accent-foreground": "143.64 89.19% 92.75%",
        "--destructive": "0 84.2% 60.2%",
        "--destructive-foreground": "210 40% 98%",
        "--border": "144.00 84.91% 10.39%",
        "--input": "143.48 22.33% 20.20%",
        "--ring": "144.00 91.09% 51.57%",
        "--chart-1": "144.00 91.09% 51.57%",
        "--chart-2": "144.19 85.65% 43.73%",
        "--chart-3": "144.08 85.96% 33.53%",
        "--chart-4": "144.19 86.11% 28.24%",
        "--chart-5": "143.76 84.87% 23.33%"
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

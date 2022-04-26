import merge from 'lodash/merge';
import { defaultTheme } from 'react-admin';

export const myTheme = merge({}, defaultTheme, {
    palette: {
        primary: {
          main: "#4082f2", // Not far from red
        },
        secondary: {
          main: "#7b3ab3", // Not far from orange
        },
        error: {
          main: "#ff5eec", // Not far from orange
        },
        type: 'dark', // Switching the dark mode on is a single property value change.
        contrastThreshold: 3,
        tonalOffset: 0.2,
    },
    typography: {
        // Use the system font instead of the default Roboto font.
        fontFamily: ['-apple-system', 'BlinkMacSystemFont', '"Poppins"', 'Arial', 'sans-serif'].join(','),
    },
    overrides: {
        MuiButton: { // override the styles of all instances of this component
            root: { // Name of the rule
                color: 'white', // Some CSS
            },
        },
    },
  });

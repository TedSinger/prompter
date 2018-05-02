Use `prompter` as a component of your PS1 in bash. It will print your current working directory, styled to be helpful:

   - If the length of your working directory approaches half of your terminal width, earlier parts of the path are truncated
   - Symlinks are in cyan, just like `ls --color`
   - Directories with open write permissions have a green background, just like `ls --color`
   - Mountpoints and cross-device symlinks are underlined

Note that if you execute `prompter` outside of your PS1 variable, it will appear to print extra escape codes. These are necessary to compensate for a... miscommunication between bash and GNU readline.

TODO:

   - Work in shells besides bash
   - Allow styles to be configured through flags
   - Get default styles from LSCOLORS, rather than hardcoding the common LSCOLORS
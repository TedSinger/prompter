Use `$(prompter)` as a component of your PS1 in bash. It will print your current working directory, styled to be helpful:

   - `prompter` truncates path components as needed to keep the prompt short, relative to your terminal width
   - `prompter` highlights symlinks, mountpoints, and open write permissions

![A demonstration of prompter](https://user-images.githubusercontent.com/2722407/39666718-4df08ec8-5076-11e8-8923-ffab51b34696.png)

```
Usage of ./prompter:
  -bash-readline-hack
        Wraps escape codes in \x01 and \x02, so that GNU Readline understands that they have no width (default true)
  -default int
        color for normal directories (default blue) (default 34)
  -open-write int
        color open write permissions (default green background) (default 42)
  -symlink int
        color for symlinks (default cyan) (default 36)
```

TODO:
   - Get default styles from LSCOLORS, rather than hardcoding the common ones
   - Get user home directory in a general way, instead of assuming /home/username
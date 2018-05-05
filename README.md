Use `$(prompter)` as a component of your PS1 in bash. It will print your current working directory, styled to be helpful:

   - `prompter` truncates path components as needed to keep the prompt short
   - `prompter` highlights symlinks, mountpoints, and open write permissions

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
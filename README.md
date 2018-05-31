Use `$(prompter)` as a component of your PS1 in the shell. It will print your current working directory, styled to be helpful:

   - `prompter` truncates path components as needed to keep the prompt short, relative to your terminal width
   - `prompter` highlights symlinks, mountpoints, and open write permissions

![A demonstration of prompter](https://user-images.githubusercontent.com/2722407/40755345-1853c926-644c-11e8-91ed-d8f076defa89.png)

```
Usage of ./prompter:
  -bash-readline-hack
      Wraps escape codes in \x01 and \x02, so that GNU Readline understands that they have no width (default true)
  -default int
      color for normal directories (ANSI escape codes) (default 34)
  -open-write int
      color open write permissions (default 42)
  -size-buffer int
      how much extra space to leave between the end of the prompt and the middle of the screen (default 10)
  -symlink int
      color for symlinks (default 36)
```

TODO:
   - Get default styles from LSCOLORS, rather than hardcoding the common ones
   - Get user home directory in a general way, instead of assuming /home/username
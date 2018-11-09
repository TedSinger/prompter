Use `$(prompter)` as a component of your PS1 in the shell. It will print your current working directory, styled to be helpful:

   - `prompter` colorizes symlinks, underlines mountpoints, and highlights open write permissions
   - `prompter` truncates path components as needed to keep the prompt short relative to your terminal width

![A demonstration of prompter](https://user-images.githubusercontent.com/2722407/40755345-1853c926-644c-11e8-91ed-d8f076defa89.png)

```
Usage of ./prompter:
  -size-buffer int
      how much extra space to leave between the end of the prompt and the middle of the screen (default 10)
  -bash-readline-hack
      Wraps escape codes in \x01 and \x02, so that GNU Readline understands that they have no width (default true)
  -default int
      style for normal directories (ANSI escape codes) (default 34)
  -open-write int
      style for open write permissions (default 42)
  -symlink int
      style for symlinks (default 36)
```

Possible next features:
   - Identify read-only filesystems, network filesystems, FUSE
   - Identify slow filesystems and full filesystems
   - Preserve characters needed for tab-completion when truncating directories
   - Get default styles from LSCOLORS, rather than hardcoding the common ones
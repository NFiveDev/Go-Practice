## Vim extension mods:
"vim.easymotion": true,
  "vim.incsearch": true,
  "vim.useSystemClipboard": true,
  "vim.useCtrlKeys": true,
  "vim.hlsearch": true,
  "vim.insertModeKeyBindings": [
    {
      "before": ["j", "j"],
      "after": ["<Esc>"]
    }
  ],
  "vim.normalModeKeyBindingsNonRecursive": [
    {
      "before": ["<leader>", "d"],
      "after": ["d", "d"]
    },
    {
      "before": ["<C-n>"],
      "commands": [":nohl"]
    },
    {
      "before": ["K"],
      "commands": ["lineBreakInsert"],
      "silent": true
    }
  ],
  "vim.leader": "<space>",
  "vim.handleKeys": {
    "<C-a>": false,
    "<C-f>": false,
    // VS Code new marker @ next occurrence
    "<C-d>": false,
    // Copy
    "<C-c>": false,
    // Cut
    "<C-x>": false,
    // Paste
    "<C-v>": false,
    "<C-z>": false,
    "<C-y>": false,
    "<C-p>": false
  },
  // To improve performance
  "extensions.experimental.affinity": {
    "vscodevim.vim": 1
  }

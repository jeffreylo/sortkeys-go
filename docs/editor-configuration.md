# Configuring your editor

## emacs

Assuming you have `sortkeys-go` in your $PATH:

```
(defun jlo/go-sortkeys ()
  (interactive)
  (message "sortkeys-go -file" (buffer-file-name))
  (shell-command (concat "sortkeys-go -w -file=" (shell-quote-argument buffer-file-name))))

(defun jlo/go-sortkeys-with-revert ()
  (interactive)
  (jlo/go-sortkeys)
  (revert-buffer t t))

(add-hook 'after-save-hook #'jlo/go-sortkeys-with-revert nil t)
```

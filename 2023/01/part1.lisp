(defun getdigits (s)
    (remove-if-not #'digit-char-p s))

(defun getnum (s)
  (parse-integer
    (format nil "~c~c" (elt s 0) (elt s (- (length s) 1)))))

(defvar input (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(format t "~d" (apply #'+ (mapcar (lambda (row) (getnum (getdigits row))) input)))

(defvar input (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defvar nums '(
    ("one" . 1)
    ("two" . 2)
    ("three" . 3)
    ("four" . 4)
    ("five" . 5)
    ("six" . 6)
    ("seven" . 7)
    ("eight" . 8)
    ("nine" . 9)))

(defun getnums (text)
  (loop for num in nums
    when (and (>= (length text) (length (car num)))
        (equal (car num) (subseq text 0 (length (car num)))))
    return (cdr num)))

(defun getdigits (text)
  (loop for x from 0 to (- (length text) 1)
    collect (if (digit-char-p (elt text x))
      (parse-integer (format nil "~c" (elt text x)))
      (getnums (subseq text x)))))

(defun parse (text)
  (let ((digits (remove nil (getdigits text))))
    (+ (* (car digits) 10) (car (last digits)))))

(print (apply #'+ (loop for text in input
    collect (parse text))))

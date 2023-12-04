(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defun parse-entry (entry)
  (let ((parts (str:split " | " entry)))
    (list (parse-card (first parts)) (parse-nums (second parts)))))

(defun parse-card (card)
  (let ((parts (str:split ": " card)))
    (parse-nums (second parts))))

(defun parse-nums (nums)
  (let ((parts (str:split " " nums :omit-nulls T)))
    (mapcar 'parse-integer parts)))

(defvar *entries* (mapcar 'parse-entry *input*))

(defun count-wins (entry)
  (length (remove nil (mapcar (lambda (num) (check-win num (car entry))) (second entry)))))

(defun check-win (num card)
  (find num card))

(print (apply '+ (loop for entry in *entries* collect
  (let ((wins (count-wins entry)))
    (if (> wins 0) (expt 2 (- wins 1)) 0)))))

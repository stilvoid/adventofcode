(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defun parse-nums (nums)
  (let ((parts (str:split " " nums :omit-nulls T)))
    (mapcar 'parse-integer parts)))

(defun parse-card (card)
  (let ((parts (str:split ": " card)))
    (parse-nums (second parts))))

(defun parse-entry (entry)
  (let ((parts (str:split " | " entry)))
    (list (parse-card (first parts)) (parse-nums (second parts)) 1)))

(defvar *entries* (mapcar 'parse-entry *input*))

(defun check-win (num card)
  (find num card))

(defun count-wins (entry)
  (length (remove nil (mapcar (lambda (num) (check-win num (car entry))) (second entry)))))

(loop for entry in *entries* for i from 1 do
  (let ((wins (count-wins entry)) (mult (third entry)))
    (loop for winner from i to (+ i wins -1) do
      (incf (third (elt *entries* winner)) mult))))

(print (apply '+ (mapcar 'third *entries*)))
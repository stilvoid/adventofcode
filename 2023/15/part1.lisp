(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (str:split "," (first *input*)))

(print *input*)

(defun hash (s)
  (let ((cur 0))
    (loop for c across s do
          (setf cur (mod (* (+ cur (char-code c)) 17) 256)))
    cur))

(print (reduce '+ (mapcar 'hash *input*)))

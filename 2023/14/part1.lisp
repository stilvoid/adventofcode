(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (loop for line in *input* collect (loop for c across line collect c)))

(defun get-pos (pos)
  (let ((x (getf pos :x)) (y (getf pos :y)))
  (when (and (>= x 0) (< x (length (first *input*))) (>= y 0) (< y (length *input*)))
    (elt (elt *input* y) x))))

(defun set-pos (pos c)
  (let ((x (getf pos :x)) (y (getf pos :y)))
  (when (and (>= x 0) (< x (length (first *input*))) (>= y 0) (< y (length *input*)))
    (setf (elt (elt *input* y) x) c)
    c)))

(defun is-free (pos)
  (let ((c (get-pos pos)))
    (eq c #\.)))

(defun move-rock-north (pos)
  (let ((new-pos (list :x (getf pos :x) :y (- (getf pos :y) 1))))
    (when (is-free new-pos)
      (set-pos pos #\.)
      (set-pos new-pos #\O)
      new-pos)))

(defun roll-rock-north (pos)
  (loop while pos do
        (setf pos (move-rock-north pos))))

(loop for y from 0 below (length *input*) do
      (loop for x from 0 below (length (first *input*)) do
            (let ((pos (list :x x :y y)))
              (when (eq (get-pos pos) #\O) (roll-rock-north pos)))))

(mapcar 'print *input*)

(print (reduce '+ (loop for y from 0 below (length *input*) collect
      (reduce '+ (remove nil (loop for x from 0 below (length (first *input*)) collect
            (let ((pos (list :x x :y y)))
              (when (eq (get-pos pos) #\O) (- (length *input*) y)))))))))

(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (loop for line in *input* for y from 0 collect
    (loop for c across line for x from 0 collect
          (list :sym c :dist nil :x x :y y :enclosed 0))))

(defconstant h (length *input*))
(defconstant w (length (first *input*)))

; Now do some processing
(loop for row in *input* do
      (loop for tile in row do
          (let ((sym (getf tile :sym)))
            (cond
              ((char= sym #\S) (defvar *start* (list :x (getf tile :x) :y (getf tile :y))))
              ((char= sym #\|) (progn (setf (getf tile :up) T) (setf (getf tile :down) T)))
              ((char= sym #\-) (progn (setf (getf tile :left) T) (setf (getf tile :right) T)))
              ((char= sym #\L) (progn (setf (getf tile :down) T) (setf (getf tile :left) T)))
              ((char= sym #\J) (progn (setf (getf tile :down) T) (setf (getf tile :right) T)))
              ((char= sym #\7) (progn (setf (getf tile :right) T) (setf (getf tile :up) T)))
              ((char= sym #\F) (progn (setf (getf tile :left) T) (setf (getf tile :up) T)))))))

(defun get-tile (pos)
  (let ((x (getf pos :x)) (y (getf pos :y)))
    (when (and (>= x 0) (>= y 0) (< x w) (< y h))
      (elt (elt *input* (getf pos :y)) (getf pos :x)))))

; Find start opts
(defvar opts nil)

(defun move-from (pos dir)
  (let ((x (getf pos :x)) (y (getf pos :y)))
    (cond
      ((eq dir :up) (list :x x :y (- y 1)))
      ((eq dir :down) (list :x x :y (+ y 1)))
      ((eq dir :left) (list :x (- x 1) :y y))
      ((eq dir :right) (list :x (+ x 1) :y y)))))

(when (getf (get-tile (move-from *start* :up)) :up) (push :up opts))
(when (getf (get-tile (move-from *start* :down)) :down) (push :down opts))
(when (getf (get-tile (move-from *start* :left)) :left) (push :left opts))
(when (getf (get-tile (move-from *start* :right)) :right) (push :right opts))

(defun opposite (motion)
  (cond
    ((eq motion :up) :down)
    ((eq motion :down) :up)
    ((eq motion :left) :right)
    ((eq motion :right) :left)))

(defun next-move (tile motion)
  (loop for opt in (list :up :down :left :right) do
    (when (and (getf tile opt) (not (eq opt motion))) (return (opposite opt)))))

(defun run-around (tile motion)
  (loop do
    (setf tile (get-tile (move-from tile motion)))
    (when (char= (getf tile :sym) #\S) (return))
    (setf (getf tile :loop) T)
    (setf motion (next-move tile motion))))

(mapcar (lambda (opt) (run-around *start* opt)) opts)

; Set the start tile properly
(setf (getf (get-tile *start*) :sym)
      (let ((opt NIL) (start-tile (get-tile *start*)))
        (setf (getf opt (first opts)) T)
        (setf (getf opt (second opts)) T)
        (setf (getf start-tile (opposite (first opts))) T)
        (setf (getf start-tile (opposite (second opts))) T)
        (setf (getf start-tile :loop) T)
        (if (getf opt :up)
          (if (getf opt :left) #\J #\L)
          (if (getf opt :left) #\7 #\F))))

(defun check-enclosure (tile motion)
  (let ((in NIL))
    (loop do
          (when (and (getf tile :loop) (getf tile :up)) (setf in (not in)))
          (when (and in (not (getf tile :loop))) (setf (getf tile :enclosed) (+ (getf tile :enclosed) 1)))
          (setf tile (get-tile (move-from tile motion)))
          (when (null tile) (return)))))

(loop for row in *input* do
    (check-enclosure (first row) :right))

(print (reduce '+ (mapcar (lambda (row) (reduce '+ (mapcar (lambda (tile) (getf tile :enclosed)) row))) *input*)))

(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(mapcar 'print *input*)

(defvar *seen* (loop for line in *input* collect
                     (loop for c across line collect 0)))

(defvar *cache* NIL)

(defconstant +height+ (length *input*))
(defconstant +width+ (length (first *input*)))

(defconstant +up+ (list :x 0 :y -1))
(defconstant +down+ (list :x 0 :y 1))
(defconstant +left+ (list :x -1 :y 0))
(defconstant +right+ (list :x 1 :y 0))

(defvar *beams* (list (list :pos (list :x 0 :y 0) :dir +right+)))

(defun pos+ (a b)
  (list :x (+ (getf a :x) (getf b :x)) :y (+ (getf a :y) (getf b :y))))

(defun do-dot (beam)
  (let ((pos (getf beam :pos))
        (dir (getf beam :dir)))
  (list :pos (pos+ pos dir) :dir dir)))

(defun do-pipe (beam)
  (let ((pos (getf beam :pos))
        (dir (getf beam :dir)))
    (if (or (equal dir +up+) (equal dir +down+))
      (list :pos (pos+ pos dir) :dir dir)
      (progn
        (push (list :pos (pos+ pos +up+) :dir +up+) *beams*)
        (list :pos (pos+ pos +down+) :dir +down+)))))

(defun do-dash (beam)
  (let ((pos (getf beam :pos))
        (dir (getf beam :dir)))
    (if (or (equal dir +left+) (equal dir +right+))
      (list :pos (pos+ pos dir) :dir dir)
      (progn
        (push (list :pos (pos+ pos +left+) :dir +left+) *beams*)
        (list :pos (pos+ pos +right+) :dir +right+)))))

(defun do-left (beam)
  (let ((pos (getf beam :pos))
        (dir (getf beam :dir)))
    (cond
      ((equal dir +up+) (list :pos (pos+ pos +left+) :dir +left+))
      ((equal dir +down+) (list :pos (pos+ pos +right+) :dir +right+))
      ((equal dir +left+) (list :pos (pos+ pos +up+) :dir +up+))
      ((equal dir +right+) (list :pos (pos+ pos +down+) :dir +down+)))))

(defun do-right (beam)
  (let ((pos (getf beam :pos))
        (dir (getf beam :dir)))
    (cond
      ((equal dir +up+) (list :pos (pos+ pos +right+) :dir +right+))
      ((equal dir +down+) (list :pos (pos+ pos +left+) :dir +left+))
      ((equal dir +left+) (list :pos (pos+ pos +down+) :dir +down+))
      ((equal dir +right+) (list :pos (pos+ pos +up+) :dir +up+)))))

(defun move (beam)
  (let* ((pos (getf beam :pos))
         (x (getf pos :x))
         (y (getf pos :y))
         (c (char (elt *input* y) x)))
    (cond
      ((char= c #\.) (do-dot beam))
      ((char= c #\|) (do-pipe beam))
      ((char= c #\-) (do-dash beam))
      ((char= c #\\) (do-left beam))
      ((char= c #\/) (do-right beam)))))

(defun do-beams ()
  (loop
    (when (= (length (remove nil (mapcar (lambda (beam) (getf beam :dead)) *beams*))) (length *beams*)) (return))
    (mapcar 'print *seen*)
    (format T "~%")
    (print (list :beams (length *beams*)))
    (loop for beam in *beams* for i from 0 do
          (let* ((beam (elt *beams* i))
                 (pos (getf beam :pos))
                 (x (getf pos :x))
                 (y (getf pos :y)))
            (when (position beam *cache* :test 'equal) (setf (getf beam :dead) T))
            (when (or (< x 0) (< y 0) (>= x +width+) (>= y +height+)) (setf (getf beam :dead) T))
            (when (not (getf beam :dead))
              (push beam *cache*)
              (setf (elt (elt *seen* y) x) (+ (elt (elt *seen* y) x) 1))
              (setf (elt *beams* i) (move beam)))))))

(do-beams)

(format T "~%")
(mapcar 'print *seen*)

(print (reduce '+ (loop for row in *seen* collect
             (count T (loop for n in row collect
                   (> n 0))))))

(print (length *cache*))

(print (list +width+ +height+ (* +width+ +height+)))

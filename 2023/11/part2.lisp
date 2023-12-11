(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (loop for line in *input* collect
                    (loop for c across line collect c)))

; Determine the size of space
(defvar *space* (loop for line in *input* collect
                      (let ((h 1))
                        (when (apply 'char= (append (list #\.) line)) (setf h 1000000))
                        (loop for i from 0 to (- (length line) 1) collect
                              (let ((w 1))
                              (let ((col (mapcar (lambda (u) (elt u i)) *input*)))
                                (when (apply 'char= (append (list #\.) col)) (setf w 1000000))
                                (list w h)))))))

(loop for line in *space* do (print line))

(print "")

(defvar *galaxies* (apply 'append (loop for line in *input* for y from 0 collect
                         (remove nil (loop for c in line for x from 0 collect
                               (when (char= c #\#) (list x y)))))))

(print *galaxies*)

(loop for galaxy in *galaxies* do
      (setf galaxy "potato"))

(defun resolve (galaxy)
  (let ((gx 0) (gy 0))
    (loop for x from 0 to (- (first galaxy) 1) do (setf gx (+ gx (first (elt (elt *space* (second galaxy)) x)))))
    (loop for y from 0 to (- (second galaxy) 1) do (setf gy (+ gy (second (elt (elt *space* y) (first galaxy))))))
    (list gx gy)))

(setf *galaxies* (mapcar 'resolve *galaxies*))

(print *galaxies*)

(print (apply '+ (loop for g1 in *galaxies* for i from 0 collect
             (apply '+ (loop for j from (+ i 1) to (- (length *galaxies*) 1) collect
                   (let ((g2 (elt *galaxies* j)))
                     (+ (abs (- (first g1) (first g2))) (abs (- (second g1) (second g2))))))))))

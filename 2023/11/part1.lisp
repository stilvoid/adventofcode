(ql:quickload "str")

(defvar *input*
  (with-open-file (file "input")
    (loop for line = (read-line file nil)
          while line
          collect line)))

(setf *input* (loop for line in *input* collect
                    (loop for c across line collect c)))

(defvar *universe* nil)

; Expand horizontal space
(loop for line in *input* do
    (when (apply 'char= (append (list #\.) line)) (push line *universe*))
    (push line *universe*))

(setf *universe* (reverse *universe*))

(loop for line in *input* do (print (coerce line 'string)))

(print "")

(loop for line in *universe* do (print (coerce line 'string)))

; Expand vertical space
(defvar *cols* (loop for i from 0 to (- (length (first *universe*)) 1) collect
      (let ((col (mapcar (lambda (line) (elt line i)) *universe*)))
        (when (apply 'char= (append (list #\.) col)) T))))

(print *cols*)

(setf *universe* (loop for line in *universe* collect
      (let ((out nil))
          (loop for c in line for i from 0 do
                (when (elt *cols* i) (push c out))
                (push c out))
          (reverse out))))

(print "")

(loop for line in *universe* do (print (coerce line 'string)))

(print "")

(setf *galaxies* (apply 'append (loop for line in *universe* for y from 0 collect
    (remove nil (loop for c in line for x from 0 collect
      (when (char= c #\#) (list x y)))))))

(print *galaxies*)

(print "")

(print (reduce '+ (loop for g1 in *galaxies* for i from 0 collect
      (reduce '+ (remove nil (loop for j from (+ i 1) to (- (length *galaxies*) 1) collect
                                   (let ((g2 (elt *galaxies* j)))
            (when (not (= i j)) (+ (abs (- (first g1) (first g2))) (abs (- (second g1) (second g2))))))))))))

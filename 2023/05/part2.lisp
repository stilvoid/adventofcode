(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defvar *seeds* (mapcar 'parse-integer (str:split " " (second (str:split ": " (first *input*))))))

(setf *seeds* (loop for i from 0 to (- (length *seeds*) 1) by 2 collect
  (let ((start (nth i *seeds*)) (len (nth (+ i 1) *seeds*)))
    (list start (+ start len -1)))))

(defun parse-map (seq)
  (loop for line in (subseq seq 1) until (string= line "") collect
    (let ((cur (mapcar 'parse-integer (str:split " " line))))
      (list (list (first cur) (+ (first cur) (third cur) -1)) (list (second cur)(+ (second cur) (third cur) -1))))))

(defun parse-maps (seq)
  (let ((out ()))
    (loop until (= (length seq) 0) collect
      (let ((next (parse-map seq)))
        (push next out)
        (if (> (length seq) (+ (length next) 1))
          (setf seq (subseq seq (+ 2 (length next))))
          (setf seq ()))))
    (nreverse out)))

(defvar *mappers* (parse-maps (subseq *input* 2)))

(defun get-overlap (range1 range2)
  (let ((start1 (first range1))
        (end1 (second range1))
        (start2 (first range2))
        (end2 (second range2)))
    (when (and (<= start2 end1) (<= start1 end2))
      (list (max start1 start2) (min end1 end2)))))

(defun get-overlaps (mapper range)
  (remove nil (loop for entry in mapper collect
    (let ((dest (first entry))
        (src (second entry))
        (overlap (get-overlap range (second entry))))
      (when overlap
        (let ((start-offset (- (first overlap) (first src))) (end-offset (- (second src) (second overlap))))
          (list overlap (list (+ (first dest) start-offset) (- (second dest) end-offset)))))))))

(defun fill-gaps (overlaps range)
  (setf overlaps (sort overlaps '< :key 'caar))
  (let ((first-overlap (first overlaps)) (last-overlap (car (last overlaps))))
    (remove nil (append
      ; start
      (when (< (first range) (first (first first-overlap))) (list (list (first range) (- (first (first first-overlap)) 1))))
      ; middle TODO (but didn't need)
      (mapcar 'second overlaps)
      ; end
      (when (> (second range) (second (first last-overlap))) (list (list (+ (second (first last-overlap)) 1) (second range))))))))

(defun map-to-next (mapper range)
  (let ((overlaps (get-overlaps mapper range)))
    (if overlaps (fill-gaps overlaps range) (list range))))

(defun map-through (range)
  (let ((cur (list range)) (next ()))
    (loop for mapper in *mappers* do
      (setf next (apply 'append (mapcar (lambda (c) (map-to-next mapper c)) cur)))
      (setf cur next))
    cur))

(print (apply 'min (mapcar (lambda (seed) (apply 'min (mapcar (lambda (result) (apply 'min result)) (map-through seed)))) *seeds*)))

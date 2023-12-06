(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defvar *times* (mapcar 'parse-integer (remove "" (str:split " " (second (str:split ":" (first *input*)))) :test 'string=)))
(defvar *distances* (mapcar 'parse-integer (remove "" (str:split " " (second (str:split ":" (second *input*)))) :test 'string=)))

(defvar *races* (loop for time in *times* for distance in *distances* collect
    (list :time time :target distance)))

(defun get-options (race-time)
  (loop for i from 1 to (- race-time 1) collect
    (* (- race-time i) i)))

(defun get-successes (race)
  (length (remove-if-not (lambda (result) (> result (getf race :target))) (get-options (getf race :time)))))

(print (apply '* (mapcar 'get-successes *races*)))

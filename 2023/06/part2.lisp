(ql:quickload "str")

(defvar *input* (with-open-file (file "input")
  (loop for line = (read-line file nil)
        while line
        collect line)))

(defvar *time* (parse-integer (apply 'str:concat (str:split " " (second (str:split ":" (first *input*)))))))
(defvar *target* (parse-integer (apply 'str:concat (str:split " " (second (str:split ":" (second *input*)))))))
;(defvar *distance* (parse-integer (apply 'str:concat (str:split " " (second (str:split ":" (second *input*)))) :test 'string=)))

(print (list *time* *target*))

(defvar *first* (loop for i from 1 to (- *time* 1) do
  (let ((result (* (- *time* i) i)))
    (when (> result *target*) (return i)))))

(print *first*)

(print (- *time* (* *first* 2) -1))

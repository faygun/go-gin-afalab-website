var gulp = require('gulp');
var less = require('gulp-less');
var livereload = require('gulp-livereload');

gulp.task('less', function() {
  gulp.src('less/app.less')
    .pipe(less())
    .pipe(gulp.dest('public/css'))
    .pipe(livereload());
});

gulp.task('watch', function() {
  livereload.listen({start:true});
  gulp.watch('less/*.less', ['less']);
});

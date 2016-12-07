module.exports = function(grunt) {

  // time the grunt execution time
  require('time-grunt')(grunt);

  // configure main project settings
  grunt.initConfig({

    // basic config
    pkg: grunt.file.readJSON('package.json'),

    // SASS compiling
    sass: {
      compile: {
        options: {
          sourcemap: 'none',
          style: 'compact'
        },
        files: [{
          expand: true,
          cwd: 'resources/assets/sass',
          src: ['*.sass'],
          dest: 'public/css',
          ext: '.full.css'
        }]
      }
    },

    // CoffeeScript compiling
    coffee: {
      compile: {
        expand: true,
        flatten: true,
        cwd: 'resources/assets/coffee',
        src: ['*.coffee'],
        dest: 'public/js',
        ext: '.full.js'
      }
    },

    // Minify css files
    cssmin: {
      minify: {
        files: [{
          expand: true,
          cwd: 'public/css',
          src: ['*.full.css'],
          dest: 'public/css',
          ext: '.css',
        }]
      }
    },

    // Minify js files
    uglify: {
      minify: {
        files: [{
          expand: true,
          cwd: 'public/js',
          src: ['*.full.js'],
          dest: 'public/js',
          ext: '.js',
        }]
      }
    },

    // watch sass and CoffeeScript files for live reloading
    watch: {
      scripts: {
        files: [
          'resources/assets/coffee/*.coffee',
          'resources/assets/sass/*.sass',
          'resources/assets/sass/*.scss',
        ],
        tasks: ['sass', 'coffee', 'cssmin', 'uglify'],
        options: {
          spawn: false
        }
      }
    }

  });

  // load plugins
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // do the tasks
  grunt.registerTask('default', ['sass', 'coffee', 'cssmin', 'uglify', 'watch']);

};

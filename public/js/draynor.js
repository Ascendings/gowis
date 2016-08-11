(function() {
  $(document).ready(function() {
    $('#menu-toggle').click(function(e) {
      e.preventDefault();
      $('#wrapper').toggleClass('menuDisplayed');
    });

    $('table').each(function() {
      if (!$(this).hasClass('table')) {
        $(this).addClass('table table-striped table-hover');
      }
    });
  });

}).call(this);

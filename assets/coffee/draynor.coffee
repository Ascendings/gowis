$(document).ready ->

  # by default, expand the sidebar if the screen is at least 1024 pixels wide
  if $(window).width() >= 1024
    $('#wrapper').toggleClass 'menuDisplayed'

  # toggle sidebar menu
  $('#menu-toggle').click (e) ->
    # prevent browser from following the link
    e.preventDefault()

    # toggle class
    $('#wrapper').toggleClass 'menuDisplayed'
    return

  # add bootstrap classess to table
  $('table').each ->
    if !$(this).hasClass('table')
      $(this).addClass 'table table-striped table-hover'
    return

  return

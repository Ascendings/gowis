$(document).ready ->

  # toggle sidebar menu
  $('#menu-toggle').click (e) ->
    # prevent browser from following the link
    e.preventDefault()

    # toggle class
    $('#wrapper').toggleClass 'menuDisplayed'
    return
    
  return

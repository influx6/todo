// Avoid `console` errors in browsers that lack a console.
(function() {
    var method;
    var noop = function () {};
    var methods = [
        'assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error',
        'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log',
        'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd',
        'timeline', 'timelineEnd', 'timeStamp', 'trace', 'warn'
    ];
    var length = methods.length;
    var console = (window.console = window.console || {});

    while (length--) {
        method = methods[length];

        // Only stub undefined methods.
        if (!console[method]) {
            console[method] = noop;
        }
    }
}());

// Place any jQuery/helper plugins in here.


//Adding placeholders
function PrettyForm(tag,submit){
  $(function() {
      $(tag).focus(function() {
          var input = $(this);
          if (input.val() == input.attr(tag)) {
              input.val('');
              input.removeClass(tag);
          }
      }).blur(function() {
          var input = $(this);
          if (input.val() == '' || input.val() == input.attr(tag)) {
              input.addClass(tag);
              input.val(input.attr(tag));
          }
      }).blur();
  });
}

function FormSubmit(tag,submit){
  $(function() {
      $(tag).submit(function(ev) {
          // ev.preventDefault()
          $(this).find(tag).each(function() {
              var input = $(this);
              if (input.val() == input.attr(tag)) {
                  input.val('');
              }
          })
          if(typeof(submit) == 'function'){
            submit(ev)
          }
      });
  });
}

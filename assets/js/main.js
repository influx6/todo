$(function(){

  PrettyForm('[placeholder]')

  var todoForm = $('.todo_new .todo_form_container')

  $(".godo_header a.new").on("click",function(e){
    e.preventDefault()
    todoForm.toggle(300)
  })

  FormSubmit('.todo_new_form',function(event){

    var scope = $(this)

    task = $(".form_inputs .task input");
    desc = $(".form_inputs .desc textarea");

    if(task.val() == ""){
      task.addClass("errorInput")
      task.attr('placeholder',"Tasks need name as well")
      event.preventDefault()
      return
    }

    task.removeClass("errorInput")

    if(desc.val() == ""){
      desc.addClass("errorInput")
      desc.attr('placeholder',"tell us a bit more about your task")
      event.preventDefault()
      return
    }

    desc.removeClass("errorInput")
  })

  $('a.delete').on("click",function(event){
    event.preventDefault()
    var id = $(this).data("id")
    $.ajax({
      url:("/todo/"+id),
      method: "delete",
      context: $(this).parent().parent(),
    }).done(function(){
      $(this).fadeOut(300).remove()
    }).error(function(){
    })

    return false
  })

  $('a.mark').on("click",function(event){
    event.preventDefault()
    var id = $(this).data("id")
    var state = $(this).data("state")

    $.ajax({
      url:("/todo/"+id),
      method: (state == "0" ? "post" : "put"),
      context: $(this).parent().parent(),
    }).done(function(){
      mark = $(this).find('a.mark')
      // content = $(this).find("div.content")
      content = $(this)

      if(state == "0"){
        mark.data("state","1")
        content.addClass('completed')
      }else{
        mark.data("state","0")
        content.removeClass('completed')
      }

    }).error(function(){
    })

    return false
  })
})

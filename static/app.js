//////////////////////////////////////////////////////////////////
///////////////////////////// homepage ///////////////////////////
//////////////////////////////////////////////////////////////////
window.onload = function(){
  if (document.getElementById('homepage-wrapper')) {
    var form = document.getElementById('sign-in-form');
    if (form) {
      document.getElementById('sign-in-form').addEventListener('submit', function(ev){
      ev.preventDefault();
      var xhr = new XMLHttpRequest();
      xhr.open("POST", '/sign-in', true);
      xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhr.onreadystatechange = function() {//Call a function when the state changes.
        if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
          console.log(xhr.responseText);
          if (xhr.responseText.length > 0) {
            location.reload();
            document.cookie = "hn_auth_token="+xhr.responseText
          } else {
            alert("incorrect password");
          }
        }
      }
      var email = document.getElementById('email').value,
      password = document.getElementById('password').value;
      xhr.send("email="+email+'&password='+password);         
      });
    }
  
    var buttons = document.getElementsByClassName('add-article');
    for (let i = 0; i < buttons.length; i++) {
      buttons[i].addEventListener('click', function(){
        var name = encodeURIComponent(this.parentElement.getElementsByClassName('title')[0].innerHTML),
        author = encodeURIComponent(this.parentElement.getElementsByClassName('author')[0].innerHTML),
        website = encodeURIComponent(this.parentElement.getElementsByClassName('website')[0].innerHTML),
        userID = document.getElementById('user-id').getAttribute('value'),
        url = encodeURIComponent(this.parentElement.getElementsByClassName('url')[0].getAttribute('href'));
        fetch('http://localhost:5050/save-article?user_id='+userID+'&author='+author+'&name='+name+'&website='+website+'&url='+url)
        .then(function(response) {
          console.log(response.status);
        });
      });
    };
  
    var signOut = document.getElementById('sign-out');
    if (signOut) {
      console.log("current user id: " +document.getElementById('user-id').getAttribute('value'));
      signOut.addEventListener('click', function(){
        document.cookie = 'hn_auth_token=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
        location.reload();
      });        
    }
  }



  var conn,
  msg = document.getElementById('submit-message'),
  log = document.getElementById('all-messages');

  function appendToChat(msg) {
    log.appendChild(msg)
  }
  var submitMessageForm = document.getElementById('submit-message-form');
  if (submitMessageForm) {
    document.getElementById('submit-message-form').addEventListener('submit', function(e){
      e.preventDefault();
        if (!conn) {
          console.log("no connection")
          return false;
        }
        if (!msg.value) {
          console.log("no value")
          return false;
        }
        conn.send(msg.value);
        msg.value = "";
        
        var messages = document.getElementById('all-messages')
        if(messages) {
          messages.scrollTop = messages.scrollHeight;
        }
        return false
    });
  }

  if (window["WebSocket"]) {
    conn = new WebSocket("ws://localhost:5050/homepage-ws");
    conn.onclose = function(evt) {
      var comment = document.createElement('div');
      comment.innerHTML = '<b>Connection closed.<\/b>';
      appendToChat(comment)
    }
    conn.onmessage = function(evt) {
      var comment = document.createElement('div'),
      fn = document.getElementById('first-name').getAttribute('value').charAt(0)
      ln = document.getElementById('last-name').getAttribute('value').charAt(0)
      comment.innerHTML = '<b>'+fn+ln+': <\/b>'+evt.data;
      appendToChat(comment)
    }
  } else {
    var comment = document.createElement('div');
    comment.innerHTML = '<b>Your browser does not support WebSockets.<\/b>';
    appendToChat(comment)
  }
}
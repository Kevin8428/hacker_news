window.onload = function(){
  if (document.getElementById('homepage-wrapper')) {
    /////////////////////////////// homepage ///////////////////////////    
    initSignInForm();
    initGoogleSignOut();
    initSignUpForm();
    initAddArticles();
    initSignOut();
    initReadStatus();
  }


  ///////////////////////////// websockets ///////////////////////////
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
    conn = new WebSocket("wss://"+window.location.hostname.replace('www-','')+":"+window.location.port+"/homepage-ws");
    conn.onclose = function(evt) {
      var comment = document.createElement('div');
      comment.innerHTML = '<b>Connection closed.<\/b>';
      appendToChat(comment)
    }
    conn.onmessage = function(evt) {
      var comment = document.createElement('div'),
          fn = document.getElementById('first-name').getAttribute('value').charAt(0)
          ln = document.getElementById('last-name').getAttribute('value').charAt(0);
      comment.innerHTML = '<b>'+fn+ln+': <\/b>'+evt.data;
      appendToChat(comment)
    }
  } else {
    var comment = document.createElement('div');
    comment.innerHTML = '<b>Your browser does not support WebSockets.<\/b>';
    appendToChat(comment)
  }    
}

var updateHomepageArticleStatus = function(articles) {
  var hpArticles = document.getElementsByClassName('url'),
      userArticles = [];
  for (let i = 0; i < hpArticles.length; i++) {
    var hpArticle = hpArticles[i].href;
    for (key in articles) {
      if (articles[key].url === hpArticle && !userArticles.includes(hpArticle)) {
        userArticles.push(hpArticle);
      }
    }
  }
  for (let i = 0; i < userArticles.length; i++) {    
  var link = document.querySelectorAll("a[href='"+userArticles[i]+"']");
    if (link) {
      var row = link[0].parentElement,
          btn = row.querySelector('button')
          savedMsg = document.createElement('div');
      row.classList.add('saved');      
      row.removeChild(btn);
      savedMsg.className = 'saved-message';
      savedMsg.innerText = 'already saved!';
      row.appendChild(savedMsg);
    }
  }
}

var initReadStatus = function() {
  var userID = document.getElementById('user-id').getAttribute('value'),
      isLoggedIn = userID != 0;
  if (isLoggedIn) {
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200)
        updateHomepageArticleStatus(JSON.parse(xhr.responseText));        
    }
    xhr.open("GET", '/user-articles?id='+userID, true); // true for asynchronous
    xhr.send(null);
  }
}

var initSignOut = function() {
  var signOut = document.getElementById('sign-out');
  if (!signOut) {
    return;
  }
  signOut.addEventListener('click', function(){
    document.cookie = 'hn_auth_token=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
    location.reload();
  });
}

var validatePasswordsFormat = function(ev) {
  var passwordInputs = ev.target.getElementsByClassName('password')
  if (passwordInputs.length === 2) {
    if (passwordInputs[0].value === passwordInputs[1].value){
      return true;
    }
  }
  return false;
}

var initSignInForm = function() {
  var form = document.getElementById('sign-in-form');
    if (!form) {
      return
    }

  form.addEventListener('submit', function(ev){
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
    var email = document.getElementById('sign-in-email').value,
    password = document.getElementById('sign-in-password').value;
    xhr.send("email="+email+'&password='+password);         
    });
}

var initGoogleSignOut = function(){
  var googleSignOutBtn = document.getElementById('google-sign-out');
  if (!googleSignOutBtn){
    return;
  }

  googleSignOutBtn.addEventListener('click', function(){
    var auth2 = gapi.auth2.getAuthInstance();
    auth2.signOut().then(function () {
      console.log('User signed out.');
    });
  });
}

var initSignUpForm = function(){
  var form = document.getElementById('sign-up-form');
  if (!form) {
    return;
  }

  form.addEventListener('submit', function(ev){
    ev.preventDefault();
    isMatchingPasswords = validatePasswordsFormat(ev);
    if (!isMatchingPasswords){
      alert('passwords must match');
      return;
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", '/sign-up', true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() {//Call a function when the state changes.
      if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
        console.log(xhr.responseText);
        if (xhr.responseText.length > 0) {
          location.reload();
          document.cookie = "hn_auth_token="+xhr.responseText
        } else {
          alert("sign up failed");
        }
      }
    }
    var email = document.getElementById('sign-up-email').value,
        password = document.getElementById('sign-up-password').value,
        fn = document.getElementById('create-user-first-name').value,
        ln = document.getElementById('create-user-last-name').value;
    xhr.send("email="+email+'&password='+password+'&first-name='+fn+'&last-name='+ln);
  });
}

var initAddArticles = function(){
  var buttons = document.getElementsByClassName('add-article');
  if (!buttons) {
    return;
  }
  for (let i = 0; i < buttons.length; i++) {
    buttons[i].addEventListener('click', function(){
      var name = encodeURIComponent(this.parentElement.getElementsByClassName('title')[0].innerHTML),
          website = encodeURIComponent(this.parentElement.getElementsByClassName('website')[0].innerHTML),
          userID = document.getElementById('user-id').getAttribute('value'),
          url = encodeURIComponent(this.parentElement.getElementsByClassName('url')[0].getAttribute('href'));
      fetch(window.location.origin+'/save-article?user_id='+userID+'&name='+name+'&website='+website+'&url='+url)
      .then(function(response) {
        console.log(response.status);
      });
    });
  };
}

var postIDToken = function(id_token) {
  console.log("preparing XHR")
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/oauth');
  xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  xhr.onload = function() {
    console.log('Signed in as: ' + xhr.responseText);
  };
  xhr.onreadystatechange = function() {
    if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200) {
      console.log("oath response: ", xhr.responseText);
      if (xhr.responseText.length > 0 && xhr.responseText != "EOF") {
        
        document.cookie = "hn_auth_token="+xhr.responseText
      } else if (xhr.responseText.length > 0 && xhr.responseText === "EOF") {
        alert("need to create a new user");
      } else {
        alert("something bad happened")
      }
    }
  }
  xhr.send('idtoken=' + id_token);
}
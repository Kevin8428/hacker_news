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
  
    var buttons = document.querySelectorAll('button');
    for (let i = 0; i < buttons.length; i++) {
      buttons[i].addEventListener('click', function(){
        var name = encodeURIComponent(this.parentElement.getElementsByClassName('name')[0].innerHTML),
        author = encodeURIComponent(this.parentElement.getElementsByClassName('author')[0].innerHTML),
        website = encodeURIComponent(this.parentElement.getElementsByClassName('website')[0].innerHTML),
        userID = document.getElementById('user-id').getAttribute('value');;
        fetch('http://localhost:5050/save-article?user_id='+userID+'&author='+author+'&name='+name+'&website='+website)
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
}
<div class="tab">
  <button class="tablinks" onclick="openTab(event, 'Go')">Go</button>
  <button class="tablinks" onclick="openTab(event, 'JavaScript')">JavaScript</button>
  <button class="tablinks" onclick="openTab(event, 'Python')">Python</button>
</div>

<div id="Go" class="tabcontent" style="display:block;">
<pre><code class="language-go">
// Go code here
fmt.Println("Hello, Go!")
</code></pre>
</div>

<div id="JavaScript" class="tabcontent" style="display:none;">
<pre><code class="language-js">
// JavaScript code here
console.log("Hello, JavaScript!");
</code></pre>
</div>

<div id="Python" class="tabcontent" style="display:none;">
<pre><code class="language-python">
# Python code here
print("Hello, Python!")
</code></pre>
</div>

<script>
  function openTab(evt, tabName) {
    var i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabcontent.length; i++) {
      tabcontent[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tablinks.length; i++) {
      tablinks[i].className = tablinks[i].className.replace(" active", "");
    }
    document.getElementById(tabName).style.display = "block";
    evt.currentTarget.className += " active";
  }
</script>

<style>
.tab {
  display: flex;
  justify-content: flex-start;
  margin-bottom: 1em;
}

.tab button {
  background-color: #sss;
  border: none;
  border-radius: 2px;
  outline: none;
  cursor: pointer;
  padding: 10px 20px;
  margin-right: 5px;
  transition: background-color 0.3s ease;
  font-size: 14px;
}

.tab button:hover {
  background-color: #03030;
}

.tab button.active {
  background-color: #ccc;
}

.tabcontent {
  display: none;
  padding: 10px 0;
  border-top: 1px solid #ccc;
}
</style>

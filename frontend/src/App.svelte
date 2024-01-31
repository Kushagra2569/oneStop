<script>
  import { Greet } from "../wailsjs/go/main/App.js";
  let val = "val";

  Greet().then((res) => {
    console.log(res);
    val = res;
  });
</script>

<main>
  <div class="flex w-screen h-screen" id="container">
    <div class="  w-20 border-r-slate-700 border-r-4" id="navbar">navbar</div>
    <div class=" bg-blue-500 w-screen" id="content">content</div>
    <div class="  w-20 border-l-slate-700 border-l-4" id="notif">notif</div>
  </div>
</main>

<style>
</style>

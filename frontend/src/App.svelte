<script>
  import { Greet } from "../wailsjs/go/main/App.js";
  import CurrentPage from "./CurrentPage.svelte";

  let currentPage = 0;

  let val = "val";

  function changePage(index) {
    currentPage = index;
  }

  Greet().then((res) => {
    val = res;
  });
</script>

<main>
  <div class="flex w-screen h-screen" id="container">
    <div
      class="  w-20 border-r-slate-700 border-r-4 flex flex-col justify-around"
      id="navbar"
    >
      <div><button on:click={() => changePage(0)}>Home</button></div>
      <div><button on:click={() => changePage(1)}>Todo</button></div>
      <div><button on:click={() => changePage(2)}>Calendar</button></div>
      <div><button on:click={() => changePage(3)}>Chat</button></div>
      <div><button on:click={() => changePage(4)}>Music</button></div>
    </div>
    <CurrentPage {currentPage} />

    <div class="  w-20 border-l-slate-700 border-l-4" id="notif">notif</div>
  </div>
</main>

<style>
</style>

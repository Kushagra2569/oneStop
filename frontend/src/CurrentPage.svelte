<script>
  import Todo from "./pages/Todo.svelte";
  import Home from "./pages/Home.svelte";
  import Calendar from "./pages/Calendar.svelte";
  import Chat from "./pages/Chat.svelte";
  import Music from "./pages/Music.svelte";
  import Weather from "./pages/Weather.svelte";

  let pages = [Home, Todo, Calendar, Chat, Music, Weather];

  export let currentPage = 0;
</script>

<main class=" bg-slate-600 w-screen" id="content">
  <svelte:component this={pages[currentPage]} />
</main>

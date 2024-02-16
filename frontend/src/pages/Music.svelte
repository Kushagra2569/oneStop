<script>
  import { OpenMultipleFiles } from "../../wailsjs/go/main/App.js";
  import { GetMusicList } from "../../wailsjs/go/main/MusicList.js";
  import { onMount } from "svelte";
  import { GetMusicListFromLocalFiles } from "../../wailsjs/go/main/MusicList.js";

  let musicList = [];

  onMount(() => {
    GetMusicList();
  });

  function openDirectory() {
    OpenMultipleFiles().then((res) => {
      if (res) {
        console.log(res);
        GetMusicListFromLocalFiles(res).then((res) => {
          musicList = res;
          console.log(musicList);
        });
      }
    });
  }
</script>

<main>
  <div class="min-h-screen flex flex-col bg-gray-100 p-4">
    <!-- Music List -->
    <div class="flex-grow overflow-auto">
      <h2 class="text-2xl mb-4">Music List</h2>
      <ul>
        <li class="mb-2 p-2 bg-white rounded shadow">
          <div class="flex justify-between items-center">
            <span>Song 1</span>
            <button class="bg-blue-500 text-white px-2 py-1 rounded"
              >Play</button
            >
          </div>
        </li>
        <!-- Repeat for each song -->
      </ul>
    </div>
    <button
      class="bg-blue-500 text-white px-2 py-1 rounded"
      on:click={openDirectory}>open file</button
    >
    <!-- Music Player -->
    <div class="mt-4 p-4 bg-white rounded shadow">
      <h2 class="text-2xl mb-4">Now Playing</h2>
      <div class="flex items-center">
        <button class="bg-blue-500 text-white px-2 py-1 rounded mr-2"
          >Prev</button
        >
        <button class="bg-blue-500 text-white px-2 py-1 rounded mr-2"
          >Play</button
        >
        <button class="bg-blue-500 text-white px-2 py-1 rounded">Next</button>
      </div>
    </div>
  </div>
</main>

<script>
  import { onMount } from "svelte";
  import { OpenMultipleFiles } from "../../wailsjs/go/main/App.js";
  import { GetMusicList } from "../../wailsjs/go/main/MusicList.js";
  import { GetMusicListFromLocalFiles } from "../../wailsjs/go/main/MusicList.js";
  import { MusicController } from "../../wailsjs/go/main/MusicList.js";

  let list;
  let musicList = [];
  let currentSong = 0;
  let audioPath = "";
  let audioElement;

  onMount(() => {
    GetMusicList().then((res) => {
      list = JSON.parse(res);
      musicList = list.musicList;
    });
  });

  function handleClickSong(index) {
    currentSong = index;
    audioPath = musicList[index].path;
    console.log("clicked on", musicList[index]);

    MusicController(index, 1).then((res) => {
      console.log(res);
    });
  }

  function playsong() {
    audioPath = musicList[currentSong].path;
    console.log("clicked on", musicList[currentSong]);

    MusicController(currentSong, 2).then((res) => {
      console.log(res);
    });
  }

  function pausesong() {
    audioPath = musicList[currentSong].path;
    console.log("clicked on", musicList[currentSong]);

    MusicController(currentSong, 0).then((res) => {
      console.log(res);
    });
  }

  function openDirectory() {
    OpenMultipleFiles().then((res) => {
      if (res) {
        GetMusicListFromLocalFiles(res).then((res) => {
          list = JSON.parse(res);
          musicList = list.musicList;
        });
      }
    });
  }
</script>

<main>
  <div class="h-screen bg-bg-color text-gray-900 text-text-color">
    <!-- Music list -->
    <div class="p-4">
      <button
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        on:click={openDirectory}>Open Directory</button
      >
      <ul>
        {#each musicList as song, index}
          <li>
            <button
              ><h class="text-white" on:click={() => handleClickSong(index)}
                >{song.title} - {song.artist}</h
              ></button
            >
          </li>
        {/each}
      </ul>
    </div>

    <!-- Play-next bar -->
    <div class="fixed bottom-0 left-0 right-0 p-4 bg-bg-color">
      <audio src={audioPath} bind:this={audioElement}></audio>
      <button
        on:click={playsong}
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >Play</button
      >
      <button
        on:click={pausesong}
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded ml-2"
        >pause</button
      >
    </div>
  </div>
</main>

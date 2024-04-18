<script>
  import { GetWeather } from "../../wailsjs/go/main/WeatherData.js";
  import { onMount } from "svelte";

  let weatherjson;
  let city = "";

  onMount(() => {
    getweatherdata();
  });

  function getweatherdata() {
    GetWeather("paris").then((res) => {
      weatherjson = JSON.parse(res);
    });
  }

  function printw() {
    console.log(weatherjson);
  }
</script>

<main>
  <div class="pt-24 flex flex-col items-center">
    <input bind:value={city} type="text" />
    <button on:click={printw}>print weather</button>
    <button on:click={getweatherdata}>get weather</button>
  </div>
  <div>
    <p bind:value={weatherjson.name}></p>
  </div>
</main>

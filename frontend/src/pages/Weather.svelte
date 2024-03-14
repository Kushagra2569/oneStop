<script>
  import { GetWeather } from "../../wailsjs/go/main/WeatherData.js";
  import { onMount } from "svelte";

  let weatherjson = "";

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
  Weather
  <button on:click={printw}>print weather</button>
</main>

<script>
  import { NewTodo } from "../../wailsjs/go/main/Todos.js";
  import { GetTodos } from "../../wailsjs/go/main/Todos.js";
  import { UpdateTodo } from "../../wailsjs/go/main/Todos.js";
  import { SaveTodos } from "../../wailsjs/go/main/Todos.js";
  import { DeleteTodo } from "../../wailsjs/go/main/Todos.js";
  import { onMount } from "svelte";

  let val = [];
  let input = "";

  onMount(() => {
    GetTodo();
  });

  function addTodo() {
    NewTodo(input, 1).then((res) => {
      //Todo kush fix priority 1 being mapped to 2
      val = JSON.parse(res);
      val = val.todos;
    });
    input = "";
  }

  function GetTodo() {
    GetTodos().then((res) => {
      val = JSON.parse(res);
      val = val.todos;
      if (!val) {
        val = [];
      }
    });
  }
  function UpdateTodoStatus(id, status, priority) {
    UpdateTodo(id, status, priority).then((res) => {
      val = JSON.parse(res);
      val = val.todos;
    });
  }

  function DeleteTodoJS(id) {
    DeleteTodo(id).then((res) => {
      val = JSON.parse(res);
      val = val.todos;
    });
  }

  function SaveTodosToFile() {
    SaveTodos();
  }
</script>

<main>
  <div>
    <input class="text-black" type="text" bind:value={input} />
    <button
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      on:click={addTodo}
      >Add
    </button>
    <button
      class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
      on:click={SaveTodosToFile}
    >
      Save
    </button>
  </div>
  <div>
    <ul>
      <p>Task , status, priority</p>
      <div class="flex items-center flex-col">
        {#each val as todo (todo.Id)}
          <li
            class="flex justify-around w-4/5 border-2 border-slate-500 rounded-lg p-4"
          >
            <p>
              {todo.Description}, Status: {todo.Status
                ? "Completed"
                : "Pending"},
              {todo.priorityValue}
            </p>
            <button
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
              on:click={() =>
                UpdateTodoStatus(todo.Id, !todo.Status, todo.priorityValue)}
              >Completed</button
            >
            <div class="relative inline-flex">
              <select
                bind:value={todo.priorityValue}
                class="border border-gray-300 rounded-full text-gray-600 h-10 pl-5 pr-10 bg-white hover:border-gray-400 focus:outline-none appearance-none"
                on:change={() =>
                  UpdateTodoStatus(todo.Id, todo.Status, todo.priorityValue)}
              >
                <option value={1}>High</option>
                <option value={2}>Medium</option>
                <option value={3}>Low</option>
              </select>
            </div>
            <button
              class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded shadow-md"
              on:click={() => DeleteTodoJS(todo.Id)}
            >
              X
            </button>
          </li>
        {/each}
      </div>
    </ul>
  </div>
</main>

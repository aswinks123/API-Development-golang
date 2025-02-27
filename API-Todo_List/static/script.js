document.getElementById("todo-form").addEventListener("submit", async function(event) {
    event.preventDefault();
    const task = document.getElementById("task").value;
    if (!task) return;
    
    const response = await fetch("/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ task, completed: false })
    });

    if (response.ok) {
        document.getElementById("task").value = "";
        fetchTasks();
    }
});

async function fetchTasks() {
    const response = await fetch("/list");
    const tasks = await response.json();
    const taskList = document.getElementById("task-list");
    taskList.innerHTML = "";
    
    tasks.forEach(task => {
        const li = document.createElement("li");
        li.textContent = task.task + (task.completed ? " (Completed)" : "");
        taskList.appendChild(li);
    });
}

// Load tasks on page load
fetchTasks();

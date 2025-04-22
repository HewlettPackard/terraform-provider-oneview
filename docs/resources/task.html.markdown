---
layout: "oneview"
page_title: "Oneview: task"
sidebar_current: "docs-task"
description: |-
 Gets information about an existing task.
---

# oneview\_task

Use this data source to access the attributes of a Task.

## Example Usage

```hcl
data "oneview_task" "test" {
 filter = "taskState='Running'"
}

resource "oneview_task" "import_task" {
 uri = "${data.oneview_task.test.uri}"
}
```

## Argument Reference

* filter - (Optional) Filter the tasks with task state is Running

* `uri` - (Required) URI of the task whose task state needs to be updated.

## Attributes Reference

* `associated_resource` - Details of associated resource.

* `associated_task_uri` - URI of another task associated with it.

* `category` - Category of this resource.

* `is_cancellable` - Boolean field indicating if the task can be cancelled or not.

* `name` - Name of the task.

* `owner` - Owner of the task.

* `parent_task_uri` - URI of the parent of the task.

* `progress_updates` - List of updates on each step being executed inside task.

* `stateReason` -  Contains the reason for changing to current state of the task.

* `task_errors` - List of errors occured during task execution.

* `task_state` - Current state of the task(Running, Completed, etc).

* `task_status` -  Short summary of the current execution.

* `taskType` - Current type of the task.

* `type` - Type of the object .

* `uri` - URI of the resource.

* `user_initiated` -  Boolean field indicating user initiated the task or not.

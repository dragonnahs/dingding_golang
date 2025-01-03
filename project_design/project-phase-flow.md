<!--
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-27 17:45:01
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-27 17:50:06
 * @FilePath: \dingding_golang\project-phase-flow.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
```mermaid
flowchart TD
    Start([开始]) --> CreateProject[创建项目]
    CreateProject --> Auto1{自动化触发}
    
    subgraph 项目创建阶段
        Auto1 -->|1.更新项目状态| UpdateStatus[更新为进行中]
        Auto1 -->|2.发送通知| NotifyManager[通知主要负责人]
    end
    
    NotifyManager --> InitPhase[发起立项阶段]
    
    subgraph 立项阶段
        InitPhase --> FillTasks[填写任务列表]
        FillTasks --> Auto2{自动化触发}
        
        Auto2 -->|遍历任务| CreateTasksRecords[创建任务记录]
        CreateTasksRecords --> NotifyExecutors[通知任务执行人]
        
        NotifyExecutors --> ParallelTasks[并行任务处理]
        ParallelTasks --> TaskProcess1[任务1处理]
        ParallelTasks --> TaskProcess2[任务2处理]
        ParallelTasks --> TaskProcessN[任务N处理]
        
        TaskProcess1 --> Auto3{自动化检查}
        TaskProcess2 --> Auto3
        TaskProcessN --> Auto3
        
        Auto3 -->|任务未完成| ParallelTasks
        Auto3 -->|全部完成| NotifyReview[通知负责人审核]
    end
    
    NotifyReview --> ManagerReview{负责人审核}
    ManagerReview -->|驳回| ParallelTasks
    ManagerReview -->|通过| Auto4{自动化触发}
    
    subgraph 进入下一阶段
        Auto4 -->|1.更新项目阶段| UpdatePhase[更新为准备阶段]
        Auto4 -->|2.发起下一阶段| NextPhase[准备阶段流程]
    end
    
    NextPhase --> End([结束])
```

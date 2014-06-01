### GOPM Registry API Specification ###

**GOPM Registry API Specification** 是 **GOPM Client** 和 **GOPM Registry** 之间交互执行的 **API** 规范

#### Ping API Specification ####

##### GET /v1/ping #####

判断服务状态

#### User API Specification ####

##### POST /v1/users/:username #####

创建用户

##### PUT /v1/users/:username/passwd ######

修改用户密码

##### POST /v1/users/:username/auth #####

登录获取 Token

##### GET /v1/users/:username #####

获取用户的基本信息

#### Repository API Specification ####

##### PUT /v1/repositories/:username/:repo #####

创建 repository 

##### GET /v1/repositories/:username/:repo/tags/:tag/json #####

获取 repository 某个 tag 的 JSON 数据

##### GET /v1/repositories/:username/:repo/tags/:tag/file #####

获取 repository 某个 tag 的二进制文件

##### PUT /v1/repositories/:username/:repo/tags/:tag/json #####

写入 repository 某个 tag 的 JSON 数据

##### PUT /v1/repositories/:username/:repo/tags/:tag/file #####

写入 repository 某个 tag 的二进制数据

##### GET /v1/repositories/:username/:repo/tags #####

获取 repository 所有 tag 的数据

##### DELETE /v1/repositories/:username/:repo #####

删除 repository 的所有数据

##### DELETE /v1/repositories/:username/:repo/tags/:tag #####

删除 repository 某个 tag 的数据及关联的二进制数据

##### GET /v1/repositories/:username/:repo

获取 repository 的所有信息

#### Search API Specification ####

##### GET /v1/search #####

搜索 repository

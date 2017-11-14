package pconst

const(
	ERROR_OK = 1001
)
const(
	ERROR_SYSTEM = 10000
)
const(
	ERROR_MONGO_SESSION =10001

	ERROR_MONGO_SEQUENCE =10002

	ERROR_MONGO_FIND = 10003

	ERROR_MONGO_INSERT = 10004

	ERROR_MONGO_ALL = 10005

	ERROR_MONGO_COUNT = 10006

	ERROR_MONGO_DISTINCT = 10007

	ERROR_MONGO_PIPE_ALL = 10008

	ERROR_MONGO_PIPE_ONE = 10009

	ERROR_MONGO_UPSERT = 10010

	ERROR_MONGO_UPDATE = 10011

	ERROR_MONGO_REMOVEID = 10012

	ERROR_MONGO_REMOVEALL = 10013
)

const(
	ERROR_MYSQL_WRITE_EMPTY= 10101

	ERROR_MYSQL_READ_EMPTY = 10102

	ERROR_MYSQL_INSERT= 10103

	ERROR_MYSQL_SELECT= 10104

	ERROR_MYSQL_UPDATE= 10105

	ERROR_MYSQL_DELETE= 10106

	ERROR_MYSQL_FIRST= 10107

	ERROR_MYSQL_COUNT= 10108
)

const(
	ERROR_CONFIG_NULL =10201

	ERRPR_CONFIG_SLICE= 10202

	ERRPR_CONFIG_SLICE_TYPE= 10203

	ERRPR_CONFIG_SLICE_CONVERT= 10204
)
const(
	ERROR_REDIS_INIT_ADDRESS=10301

	ERROR_REDIS_POOL_NULL=10302

	ERROR_REDIS_POOL_GET=10303

	ERROR_REDIS_POOL_EMPTY=10304

	ERROR_REDIS_POOL_REDIAL = 10305

	ERROR_REDIS_SET_MARSHAL = 10306

	ERROR_REDIS_SET_DO = 10307

	ERROR_REDIS_MSET_REPLY= 10308

	ERROR_REDIS_MSET_MARSHAL =10309

	ERROR_REDIS_MSET_DO =10310

	ERROR_REDIS_GET_DO =10311

	ERROR_REDIS_GET_UNMARSHAL =10312

	ERROR_REDIS_MGET_TYPE =10313

	ERROR_REDIS_MGET_DO =10314

	ERROR_REDIS_INCR_DO =10315

	ERROR_REDIS_INCR_CONVERT =10316

	ERROR_REDIS_DEL_DO =10317

	ERROR_REDIS_EXPIRE_DO =10318


	ERROR_REDIS_DO =10319

	ERROR_REDIS_CONVERT =10320

	ERROR_REDIS_PIPE_SEND =10321

	ERROR_REDIS_PIPE_FLUSH =10322

	ERROR_REDIS_PIPE_RECEIVE =10323

	ERROR_REDIS_PIPE_UNMARSHAL =10324

	ERROR_REDIS_SETNX_REPLY = 10325

	ERROR_REDIS_ZADDM_CONVERT = 10326
)

const(
	ERROR_GRPC_CONFIG = 10401

	ERROR_GRPC_DAIL = 10402

	ERROR_GRPC_INVOKE = 10403
)

const(
	ERROR_HTTP_CONFIG = 10501

	ERROR_HTTP_POSTFORM = 10502

	ERROR_HTTP_POSTFORM_RESPONSE=10503

	ERROR_HTTP_POSTGET = 10504

	ERROR_HTTP_POSTGET_RESPONSE=10505

	ERROR_HTTP_READ=10506

	ERROR_HTTP_UNMARSHAL=10507
)

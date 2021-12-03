// Softbery Libraries (c) by Paweł Tobis
// 
// Softbery Libraries is licensed under a
// Creative Commons Attribution-ShareAlike 4.0 International License.
// 
// You should have received a copy of the license along with this
// work. If not, see <http://creativecommons.org/licenses/by-sa/4.0/>.

package config/types

type ConfigEnum int64

const {
	Loaded 			 ConfigEnum = 0
	Unloaded 		 ConfigEnum = 1
	LoadedWithErrors ConfigEnum = 2
}

type Config struct {
	Key string
	Value string
	Status 
}

type Configs []Config